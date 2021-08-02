package api

import (
	"context"
	"encoding/json"
	"github.com/chancetudor/dubzone-api/internal/auth"
	"github.com/chancetudor/dubzone-api/internal/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
	"time"
)

const WeaponCollectionName = "Weapons"

/* Weapon endpoints live here
* functionality for
	* Creating a single weapon
	* Reading a single weapon
	* Updating a single weapon
	* Reading multiple weapons, given weapon name, or returning all weapons
*/

/*
	// single weapon endpoints
	r.HandleFunc("/weapon/{weaponname}", CreateWeaponEndpoint).Methods("POST")
	r.HandleFunc("/weapon/{weaponname}", ReadWeaponEndpoint).Methods("GET")
	r.HandleFunc("/weapon/{weaponname}", UpdateWeaponEndpoint).Methods("PUT")
	// multiple weapon endpoints
	r.HandleFunc("/weapons", ReadWeaponEndpoint).Methods("GET")
*/

// CreateWeaponEndpoint creates a single new weapon in the Weapons collection
func CreateWeaponEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var weapon models.Weapon
	err := json.NewDecoder(request.Body).Decode(&weapon)
	weapon.WeaponName = strings.ToUpper(weapon.WeaponName)
	// client is used to connect to MongoDB directly
	var client = NewClient()
	var db = auth.NewAuth().Database
	collection := client.Database(db).Collection(WeaponCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer client.Disconnect(ctx)
	defer cancel()
	_, err = collection.InsertOne(ctx, weapon)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.WithFields(log.Fields{
			"func":  "CreateWeaponEndpoint()",
			"event": "Inserting into collection",
		}).Error(err)
		return
	}
	err = json.NewEncoder(response).Encode(weapon.WeaponName)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.WithFields(log.Fields{
			"func":  "CreateWeaponEndpoint()",
			"event": "Encoding a weapon name as a response",
		}).Error(err)
		return
	}
	response.Write([]byte(`{"message": "Weapon added"}`))
}

// ReadWeaponEndpoint returns weapon data for a specified weapon name
func ReadWeaponEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var weapon models.Weapon
	var client = NewClient()
	var db = auth.NewAuth().Database
	var collectionName = auth.NewAuth().WeaponsCollection
	collection := client.Database(db).Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer client.Disconnect(ctx)
	defer cancel()

	// find weapon using given weaponname
	cursor, err := collection.Find(ctx, bson.M{{""}})
}

func UpdateWeaponEndpoint(response http.ResponseWriter, request *http.Request) {

}

func ReadWeaponsEndpoint(response http.ResponseWriter, request *http.Request) {

}

/*
	* we probably don't need to delete a weapon
func DeleteWeaponEndpoint(response http.ResponseWriter, request *http.Request) {

}
*/

