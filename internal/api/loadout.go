package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chancetudor/dubzone-api/internal/auth"
	"github.com/chancetudor/dubzone-api/internal/entity"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// Loadout endpoints live here TODO FINISH

// CreateLoadoutEndpoint creates a single new loadout,
// either with a weaponname variable pointing to the weapon in the DB that needs a new loadout
// or creating a new loadout independent of a weapon
func CreateLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {
	// TODO custom error handling with logging
	response.Header().Add("content-type", "application/json")
	var loadout entity.Loadout
	vars := mux.Vars(request)
	// then a loadout is being created for a specific supplied weapon name
	if len(vars) != 0 {
		// search DB for matching weapon name
		// insert a new loadout in the []Loadouts struct for that weapon
		weaponName := vars["weaponname"]
		fmt.Println(weaponName)
	}
	err := json.NewDecoder(request.Body).Decode(&loadout)
	if err != nil {
		log.Fatal(err)
	}
	// client is used to connect to MongoDB directly
	var client = NewClient()
	var db = auth.NewAuth().Database
	collection := client.Database(db).Collection("Loadouts")
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
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	response.Write([]byte(`{"message": "Weapon added"}`))
}

func ReadLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {

}

func UpdateLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {

}

func DeleteLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {

}

func ReadLoadoutsEndpoint(response http.ResponseWriter, request *http.Request) {

}

func ReadLoadoutsByCategoryEndpoint(response http.ResponseWriter, request *http.Request) {

}
