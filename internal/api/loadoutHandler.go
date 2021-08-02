package api

import (
	"context"
	"encoding/json"
	"github.com/chancetudor/dubzone-api/internal/models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
	"time"
)

/* Loadout endpoints live here
* functionality for
	* Creating a single loadout
	* Reading loadouts given a category
	* Reading loadouts given a weapon name
*/

// CreateLoadoutEndpoint creates a single new loadout in the loadouts collection
func (a *API) CreateLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {
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

	db := a.Auth.Database
	collection := a.Client.Database(db).Collection(a.Auth.LoadoutCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, loadout)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.WithFields(log.Fields{
			"func":  "CreateLoadoutEndpoint()",
			"event": "Inserting into collection",
		}).Error(err)
		return
	}

	err = json.NewEncoder(response).Encode(loadout.Weapon)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.WithFields(log.Fields{
			"func":  "CreateLoadoutEndpoint()",
			"event": "Encoding a weapon name as a response",
		}).Error(err)
		return
	}

	response.Write([]byte(`{"message": "Weapon added"}`))
}

// ReadLoadoutsEndpoint returns loadouts for
// a given category,
// a given weapon name,
// or returns all loadouts if category / weapon name are not provided
func (a *API) ReadLoadoutsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := request.URL.Query()
	// category and weapon are optional query parameters and are stored
	// in the database in Uppercase, so we capitalize the query params
	category := strings.Title(params.Get("category"))
	weapon := strings.ToUpper(params.Get("weapon"))
	var loadouts []models.Loadout

	// return loadouts for specific category
	if category != "" {
		query := bson.M{"category": category}
		loadouts = a.readManyLoadouts(query)
	} else if weapon != "" { // else return loadouts for specific weapon
		query := bson.M{"weapon": weapon}
		loadouts = a.readManyLoadouts(query)
	} else { // else return all loadouts in the loadouts collection
		query := bson.M{}
		loadouts = a.readManyLoadouts(query)
	}

	if loadouts == nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "cannot find loadouts"}`))
		log.WithFields(log.Fields{
			"func":  "ReadLoadoutsEndpoint()",
			"event": "No loadouts retrieved",
		}).Error()
		return
	}

	err := json.NewEncoder(response).Encode(loadouts)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.WithFields(log.Fields{
			"func":  "ReadLoadoutsEndpoint()",
			"event": "Encoding loadouts into JSON response",
		}).Error(err)
		return
	}
}

// readManyLoadouts is a helper function to retrieve all loadouts
// and contains the true logic for querying the database
func (a *API) readManyLoadouts(query bson.M) []models.Loadout {
	db := a.Auth.Database
	collection := a.Client.Database(db).Collection(a.Auth.LoadoutCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// find all documents using the given bson.M{} query,
	// where the bson.M{} query can specify category, weapon, or be empty (find all loadouts)
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

	var loadouts []models.Loadout
	// TODO optimize for loop
	// iterate through cursor and encode documents into Loadout struct
	for cursor.Next(ctx) {
		var loadout models.Loadout
		if err = cursor.Decode(&loadout); err != nil {
			log.WithFields(log.Fields{
				"func":  "readManyLoadouts()",
				"event": "Decoding a cursor into a Loadout",
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
	* if necessary, we just add a new loadout
func UpdateLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {

}
*/

/*
	* we probably don't need to delete loadouts
func DeleteLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {

}
*/