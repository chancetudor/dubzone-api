package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chancetudor/dubzone-api/internal/auth"
	"github.com/chancetudor/dubzone-api/internal/models"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	var loadout models.Loadout
	// decode JSON request payload into Loadout
	err := json.NewDecoder(request.Body).Decode(&loadout)
	if err != nil {
		log.WithFields(log.Fields{
			"func": "CreateLoadoutEndpoint()",
			"event": "Decoding JSON to loadout struct",
			"line": 38,
		}).Fatal(err)
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
		log.WithFields(log.Fields{
			"func": "CreateLoadoutEndpoint()",
			"event": "Inserting into collection",
			"line": 52,
		}).Error(err)
		return
	}
	err = json.NewEncoder(response).Encode(loadout.Weapon)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.WithFields(log.Fields{
			"func": "CreateLoadoutEndpoint()",
			"event": "Encoding a weapon name as a response",
			"line": 64,
		}).Error(err)
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
	var loadout models.Loadout
	vars := mux.Vars(request)
	weaponName := strings.Title(vars["weaponname"])
	err := collection.FindOne(ctx, models.Loadout{Weapon: weaponName}).Decode(&loadout)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.WithFields(log.Fields{
			"func": "ReadLoadoutEndpoint()",
			"event": "Finding a weapon in DB",
			"line": 90,
		}).Error(err)
		return
	}
	err = json.NewEncoder(response).Encode(loadout)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.WithFields(log.Fields{
			"func": "ReadLoadoutEndpoint()",
			"event": "Encoding loadout JSON response",
			"line": 101,
		}).Error(err)
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
	var loadouts []models.Loadout
	// return all loadouts for matching category
	if len(vars) != 0 {
		// search DB for matching category
		category := strings.Title(vars["category"])
		loadouts = readManyLoadouts(category)
		err := json.NewEncoder(response).Encode(loadouts)
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(`{"message": "` + err.Error() + `"}`))
			log.WithFields(log.Fields{
				"func": "ReadLoadoutsEndpoint()",
				"event": "Encoding loadouts into JSON response",
				"line": 132,
			}).Error(err)
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
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.WithFields(log.Fields{
			"func": "ReadLoadoutsEndpoint()",
			"event": "Finding all loadouts",
			"line": 150,
		}).Error(err)
		return
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)
	// TODO optimize for loop
	for cursor.Next(ctx) {
		var loadout models.Loadout
		err := cursor.Decode(&loadout)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"message": "` + err.Error() + `"}`))
			log.WithFields(log.Fields{
				"func": "ReadLoadoutsEndpoint()",
				"event": "Decoding cursor into loadouts",
				"line": 167,
			}).Error(err)
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
		log.WithFields(log.Fields{
			"func": "ReadLoadoutsEndpoint()",
			"event": "Encoding []Loadouts into JSON response",
			"line": 185,
		}).Error(err)
		return
	}
}

func readManyLoadouts(category string) []models.Loadout {
	// client is used to connect to MongoDB directly
	var client = NewClient()
	var db = auth.NewAuth().Database
	collection := client.Database(db).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.D{{"category", category}})
	if err != nil {
		log.WithFields(log.Fields{
			"func": "readManyLoadouts()",
			"event": "Finding a weapon in DB",
			"line": 205,
		}).Error(err)
		return nil
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)
	var loadouts []models.Loadout
	// TODO optimize for loop
	for cursor.Next(ctx) {
		var loadout models.Loadout
		err = cursor.Decode(&loadout)
		if err != nil {
			log.WithFields(log.Fields{
				"func": "readManyLoadouts()",
				"event": "Decoding a cursor into Loadout",
				"line": 221,
			}).Error(err)
			return nil
		}
		loadouts = append(loadouts, loadout)
	}
	if err := cursor.Err(); err != nil {
		return nil
	}

	return loadouts
}
