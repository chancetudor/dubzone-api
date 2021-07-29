package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chancetudor/dubzone-api/internal/auth"
	"github.com/chancetudor/dubzone-api/internal/entity"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strings"
	"time"
)

const CollectionName = "Loadouts"

// Loadout endpoints live here TODO FINISH

// CreateLoadoutEndpoint creates a single new loadout,
// either with a weaponname variable pointing to the weapon in the DB that needs a new loadout
// or creating a new loadout independent of a weapon
func CreateLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {
	// TODO custom error handling with logging
	response.Header().Add("content-type", "application/json")
	vars := mux.Vars(request)
	// then a loadout is being created for a specific supplied weapon name
	if len(vars) != 0 {
		// TODO separate function?
		// search DB for matching weapon name
		// insert a new loadout in the []Loadouts struct for that weapon
		weaponName := vars["weaponname"]
		fmt.Println(weaponName)
	}
	var loadout entity.Loadout
	err := json.NewDecoder(request.Body).Decode(&loadout)
	if err != nil {
		log.Fatal(err)
	}
	// client is used to connect to MongoDB directly
	var client = NewClient()
	var db = auth.NewAuth().Database
	collection := client.Database(db).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = collection.InsertOne(ctx, loadout)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	err = json.NewEncoder(response).Encode(loadout.Weapon)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	response.Write([]byte(`{"message": "Weapon added"}`))
}

func ReadLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {
	// TODO custom error handling with logging
	response.Header().Add("content-type", "application/json")
	// client is used to connect to MongoDB directly
	var client = NewClient()
	var db = auth.NewAuth().Database
	collection := client.Database(db).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var loadout entity.Loadout
	vars := mux.Vars(request)
	weaponName := strings.Title(vars["weaponname"])
	err := collection.FindOne(ctx, entity.Loadout{Weapon: weaponName}).Decode(&loadout)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	err = json.NewEncoder(response).Encode(loadout)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
}

func UpdateLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {

}

func DeleteLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {

}

func ReadLoadoutsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	vars := mux.Vars(request)
	var loadouts []entity.Loadout
	// return all loadouts for matching category
	if len(vars) != 0 {
		// TODO separate function?
		// search DB for matching category
		category := strings.Title(vars["category"])
		loadouts = readManyLoadouts(category)
		err := json.NewEncoder(response).Encode(loadouts)
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(`{"message": "` + err.Error() + `"}`))
			return
		}
		return
	}
	// client is used to connect to MongoDB directly
	var client = NewClient()
	var db = auth.NewAuth().Database
	collection := client.Database(db).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	// TODO custom error handling
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	} (cursor, ctx)
	// TODO optimize for loop
	for cursor.Next(ctx) {
		var loadout entity.Loadout
		err := cursor.Decode(&loadout)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"message": "` + err.Error() + `"}`))
			return
		}
		loadouts = append(loadouts, loadout)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	err = json.NewEncoder(response).Encode(loadouts)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
}

func readManyLoadouts(category string) []entity.Loadout {
	// client is used to connect to MongoDB directly
	var client = NewClient()
	var db = auth.NewAuth().Database
	collection := client.Database(db).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.D{{"category", category}})
	if err != nil {
		// TODO custom error handling
		return nil
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	} (cursor, ctx)
	var loadouts []entity.Loadout
	// TODO optimize for loop
	for cursor.Next(ctx) {
		var loadout entity.Loadout
		err = cursor.Decode(&loadout)
		if err != nil {
			// TODO custom error handling
			return nil
		}
		loadouts = append(loadouts, loadout)
	}
	if err := cursor.Err(); err != nil {
		// TODO custom error handling
		return nil
	}

	return loadouts
}

