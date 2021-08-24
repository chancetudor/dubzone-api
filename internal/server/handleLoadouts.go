package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chancetudor/dubzone-api/internal/logger"
	"github.com/chancetudor/dubzone-api/internal/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
	"time"
)

/* Loadout endpoints live here
* functionality for
	* Reading all loadouts
	* Reading loadouts given a category
	* Reading loadouts given a weapon name
*/

func (srv *server) CreateLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {
	var loadout models.Loadout
	if err := json.NewDecoder(request.Body).Decode(&loadout); err != nil {
		srv.respond(response, err, http.StatusInternalServerError)
		logger.Error(err, "Error decoding JSON body", "CreateLoadoutEndpoint")
		return
	}
	// put weapon name in ALL CAPS
	loadout.Primary.WeaponName = strings.ToUpper(loadout.Primary.WeaponName)
	loadout.Secondary.WeaponName = strings.ToUpper(loadout.Secondary.WeaponName)

	db := srv.Auth.Database
	collection := srv.Client.Database(db).Collection(srv.Auth.LoadoutCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, loadout)
	if err != nil {
		srv.respond(response, err, http.StatusInternalServerError)
		logger.Error(err, "Error decoding JSON body", "CreateLoadoutEndpoint")
		return
	}

	srv.respond(response, result, http.StatusOK)
}

// ReadLoadoutsEndpoint returns all loadouts
// GET /loadouts
func (srv *server) ReadLoadoutsEndpoint(response http.ResponseWriter, request *http.Request) {
	query := bson.D{}
	// projection to suppress loadout ID
	// p := bson.D{{"_id", 0}}
	loadouts, err := srv.getLoadouts(query)

	if err != nil {
		srv.respond(response, err, http.StatusInternalServerError)
		logger.Error(nil, "No loadouts retrieved", "ReadLoadoutsEndpoint")
		return
	}

	if loadouts == nil {
		srv.respond(response, nil, http.StatusNoContent)
	}

	srv.respond(response, loadouts, http.StatusOK)
}

// ReadLoadoutsByMetaEndpoint returns loadouts that are listed as meta in the DB
// GET /loadouts/meta
func (srv *server) ReadLoadoutsByMetaEndpoint(response http.ResponseWriter, request *http.Request) {
	query := bson.D{{"meta_loadout", true}}
	// projection to suppress loadout ID
	// p := bson.D{{"_id", 0}}
	loadouts, err := srv.getLoadouts(query)

	if err != nil {
		srv.respond(response, err, http.StatusInternalServerError)
		logger.Error(err, "No loadouts retrieved", "ReadLoadoutsByCategoryEndpoint")
		return
	}

	if loadouts == nil {
		srv.respond(response, nil, http.StatusNoContent)
	}

	srv.respond(response, loadouts, http.StatusOK)
}

// ReadLoadoutsByCategoryEndpoint returns all loadouts with a specified category
// GET /loadouts/category/{cat}
func (srv *server) ReadLoadoutsByCategoryEndpoint(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	// category is stored in the database in Uppercase, so we capitalize the query param
	category := strings.Title(params["cat"])
	query := bson.D{{"primary.category", category}}
	// projection to suppress loadout ID
	// p := bson.D{{"_id", 0}}
	loadouts, err := srv.getLoadouts(query)

	if err != nil {
		srv.respond(response, err, http.StatusInternalServerError)
		logger.Error(err, "No loadouts retrieved", "ReadLoadoutsByCategoryEndpoint")
		return
	}

	if loadouts == nil {
		srv.respond(response, []byte("No content"), http.StatusNoContent)
		return
	}

	srv.respond(response, loadouts, http.StatusOK)
}

// ReadLoadoutsByWeaponEndpoint returns all loadouts for a specified weapon
// GET /loadouts/weapon/{name}
func (srv *server) ReadLoadoutsByWeaponEndpoint(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	// weapon_name is stored in the database in ALL CAPS, so we UPPERCASE the query param
	weaponName := strings.ToUpper(params["name"])
	query := bson.D{{"primary.weapon_name", weaponName}}
	// projection to suppress loadout ID
	// p := bson.D{{"_id", 0}}
	loadouts, err := srv.getLoadouts(query)

	if err != nil {
		srv.respond(response, err, http.StatusInternalServerError)
		logger.Error(err, "No loadouts retrieved", "ReadLoadoutsByWeaponEndpoint")
		return
	}

	srv.respond(response, loadouts, http.StatusOK)
}

// getLoadouts is a helper function to retrieve all loadouts
// and contains the true logic for querying the database
func (srv *server) getLoadouts(query bson.D) ([]models.Loadout, error) {
	fmt.Println(query)
	db := srv.Auth.Database
	collection := srv.Client.Database(db).Collection(srv.Auth.LoadoutCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// find all documents using the given bson.M{} query,
	// where the bson.D{} query can specify category, weapon, or be empty (find all loadouts)
	cursor, err := collection.Find(ctx, query)
	if err != nil {
		logger.Error(err, "Finding loadout in DB", "getLoadouts")
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)

	var loadouts []models.Loadout
	for cursor.Next(ctx) {
		var loadout models.Loadout
		if err = cursor.Decode(&loadout); err != nil {
			logger.Error(err, "Decoding cursor into Loadout struct", "getLoadouts")
			return nil, err
		}
		// append encoded Loadout in []Loadouts
		loadouts = append(loadouts, loadout)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return loadouts, nil
}
