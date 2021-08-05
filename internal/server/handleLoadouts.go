package server

import (
	"context"
	"encoding/json"
	"github.com/chancetudor/dubzone-api/internal/models"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
	"time"
)

/* Loadout endpoints live here
* functionality for
	* Creating srv single loadout
	* Reading loadouts given srv category
	* Reading loadouts given srv weapon name
*/

// CreateLoadoutEndpoint creates single new loadout in the loadouts collection
// POST /loadout
func (srv *server) CreateLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var loadout models.Loadout
	// decode JSON request payload into Loadout
	err := json.NewDecoder(request.Body).Decode(&loadout)
	if err != nil {
		log.WithFields(log.Fields{
			"func":  "CreateLoadoutEndpoint()",
			"event": "Decoding JSON to loadout struct",
		}).Fatal(err)
	}
	// capitalize weapon name to match DB schema
	loadout.Weapon = strings.ToUpper(loadout.Weapon)

	db := srv.Auth.Database
	collection := srv.Client.Database(db).Collection(srv.Auth.LoadoutCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, loadout)

	if err != nil {
		srv.respond(response, nil, http.StatusInternalServerError)
		log.WithFields(log.Fields{
			"func":  "CreateLoadoutEndpoint()",
			"event": "Inserting into collection",
		}).Error(err)
		return
	}

	srv.respond(response, loadout.Weapon, http.StatusOK)
}

// ReadLoadoutsEndpoint returns all loadouts
// GET /loadouts
func (srv *server) ReadLoadoutsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var loadouts []models.Loadout
	query := bson.M{}
	loadouts = srv.readManyLoadouts(query)

	if loadouts == nil {
		srv.respond(response, loadouts, http.StatusInternalServerError)
		log.WithFields(log.Fields{
			"func":  "ReadLoadoutsEndpoint()",
			"event": "No loadouts retrieved",
		}).Error()
		return
	}

	srv.respond(response, loadouts, http.StatusOK)
}

// ReadLoadoutsByCategoryEndpoint returns all loadouts with srv specified category
// GET /loadouts/{category}
func (srv *server) ReadLoadoutsByCategoryEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	// category and weapon are optional query parameters and are stored
	// in the database in Uppercase, so we capitalize the query params
	category := strings.Title(params["category"])
	var loadouts []models.Loadout
	query := bson.M{"category": category}
	loadouts = srv.readManyLoadouts(query)

	if loadouts == nil {
		srv.respond(response, loadouts, http.StatusInternalServerError)
		log.WithFields(log.Fields{
			"func":  "ReadLoadoutsEndpoint()",
			"event": "No loadouts retrieved",
		}).Error()
		return
	}

	srv.respond(response, loadouts, http.StatusOK)
}

// ReadLoadoutsByWeaponEndpoint returns all loadouts for srv specified weapon
// GET /loadouts/{weaponname}
// TODO deal with weapon names containing spaces -- maybe in docs specify that spaces must be represented by "_"?
func (srv *server) ReadLoadoutsByWeaponEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	// category and weapon are optional query parameters and are stored
	// in the database in Uppercase, so we capitalize the query params
	weaponName := strings.ToUpper(params["weaponname"])
	var loadouts []models.Loadout
	query := bson.M{"category": weaponName}
	loadouts = srv.readManyLoadouts(query)

	if loadouts == nil {
		srv.respond(response, loadouts, http.StatusInternalServerError)
		log.WithFields(log.Fields{
			"func":  "ReadLoadoutsEndpoint()",
			"event": "No loadouts retrieved",
		}).Error()
		return
	}

	srv.respond(response, loadouts, http.StatusOK)
}

// readManyLoadouts is srv helper function to retrieve all loadouts
// and contains the true logic for querying the database
func (srv *server) readManyLoadouts(query bson.M) []models.Loadout {
	db := srv.Auth.Database
	collection := srv.Client.Database(db).Collection(srv.Auth.LoadoutCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// find all documents using the given bson.M{} query,
	// where the bson.M{} query can specify category, weapon, or be empty (find all loadouts)
	cursor, err := collection.Find(ctx, query)
	if err != nil {
		log.WithFields(log.Fields{
			"func":  "readManyLoadouts()",
			"event": "Finding srv loadout in DB",
		}).Error(err)
		return nil
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)

	var loadouts []models.Loadout
	// TODO optimize for loop
	// iterate through cursor and encode documents into Loadout struct
	for cursor.Next(ctx) {
		var loadout models.Loadout
		if err = cursor.Decode(&loadout); err != nil {
			log.WithFields(log.Fields{
				"func":  "readManyLoadouts()",
				"event": "Decoding srv cursor into srv Loadout",
			}).Error(err)
			return nil
		}
		// append encoded Loadout in []Loadouts
		loadouts = append(loadouts, loadout)
	}
	if err := cursor.Err(); err != nil {
		return nil
	}

	return loadouts
}

/*
	* we probably don't need to update loadouts
	* if necessary, we just add srv new loadout
func UpdateLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {

}
*/

/*
	* we probably don't need to delete loadouts
func DeleteLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {

}
*/
