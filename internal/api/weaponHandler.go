package api

import (
	"context"
	"encoding/json"
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

	client := NewClient()
	authDetails := auth.NewAuth()
	db := authDetails.Database
	collection := client.Database(db).Collection(authDetails.WeaponsCollection)
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

	client := NewClient()
	authDetails := auth.NewAuth()
	db := authDetails.Database
	collection := client.Database(db).Collection(authDetails.WeaponsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer client.Disconnect(ctx)
	defer cancel()

	params := mux.Vars(request)
	weaponName := strings.ToUpper(params["weaponname"])
	var weapon models.Weapon
	// find weapon using given weaponname
	err := collection.FindOne(ctx, bson.D{{"weapon_name", weaponName}}).Decode(&weapon)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.WithFields(log.Fields{
			"func":  "ReadWeaponEndpoint()",
			"event": "Finding weapon in database",
		}).Error(err)
		return
	}

	err = json.NewEncoder(response).Encode(weapon)
}

// ReadWeaponsEndpoint returns multiple weapons,
// either all weapons in the Weapons collection
// or all weapons from an optionally specified game
func ReadWeaponsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := request.URL.Query()
	game := strings.ToUpper(params.Get("game"))

	var weapons []models.Weapon
	if game != "" {
		query := bson.M{"game_from": game}
		weapons = readManyWeapons(query)
	} else {
		query := bson.M{}
		weapons = readManyWeapons(query)
	}

	if weapons == nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "cannot find loadouts"}`))
		log.WithFields(log.Fields{
			"func":  "ReadWeaponsEndpoint()",
			"event": "No weapons retrieved",
		}).Error()
		return
	}

	err := json.NewEncoder(response).Encode(weapons)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.WithFields(log.Fields{
			"func":  "ReadWeaponsEndpoint()",
			"event": "Encoding weapons into JSON response",
		}).Error(err)
		return
	}
}

// readManyWeapons is a helper function for ReadWeaponsEndpoint
// and contains the true logic for querying the database
func readManyWeapons(query bson.M) []models.Weapon {
	client := NewClient()
	authDetails := auth.NewAuth()
	db := authDetails.Database
	collection := client.Database(db).Collection(authDetails.WeaponsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer client.Disconnect(ctx)
	defer cancel()

	// find all documents using the given bson.M{} query,
	// where the bson.M{} query can specify the game a weapon is from
	// or the bson.M{} query can be empty (find all weapons)
	cursor, err := collection.Find(ctx, query)
	if err != nil {
		log.WithFields(log.Fields{
			"func":  "readManyLoadouts()",
			"event": "Finding a loadout in DB",
		}).Error(err)
		return nil
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)

	var weapons []models.Weapon
	// TODO optimize for loop
	// iterate through cursor and encode documents into Loadout struct
	for cursor.Next(ctx) {
		var weapon models.Weapon
		if err = cursor.Decode(&weapon); err != nil {
			log.WithFields(log.Fields{
				"func":  "readManyWeapons()",
				"event": "Decoding a cursor into a Weapon",
			}).Error(err)
			return nil
		}
		// append encoded Loadout in []Loadouts
		weapons = append(weapons, weapon)
	}
	if err := cursor.Err(); err != nil {
		return nil
	}

	return weapons
}

/*
	* we probably don't need to delete a weapon
func DeleteWeaponEndpoint(response http.ResponseWriter, request *http.Request) {

}
*/

/*
	* we probably don't need to update a weapon
// UpdateWeaponEndpoint takes a specified weapon and a series of parameters
// and updates that weapon's parameters to the given ones
func UpdateWeaponEndpoint(response http.ResponseWriter, request *http.Request) {

}
*/

