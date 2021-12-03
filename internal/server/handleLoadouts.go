package server

import (
	"context"
	"github.com/chancetudor/dubzone-api/internal/models"
	"github.com/pkg/errors"
	"net/http"
)

// LoadoutKey is the key to be used for use in middleware.
// When we store the internal representation of a JSON struct,
// we use the context.WithValue function, which takes this as its key.
type LoadoutKey struct{}

/*
CreateLoadout takes a JSON representation of a Loadout,
marshals it into a struct of type Loadout,
and TODO stores that struct in a database.
*/
func (srv *server) CreateLoadout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loadout := r.Context().Value(LoadoutKey{}).(*models.Loadout)
		models.AddProduct(loadout)
	}
}

// swagger:route GET /loadouts loadouts listLoadouts
// Returns a list of all loadouts.
// responses:
//	200: loadoutsResponse

// GetLoadouts returns all loadouts in the database in JSON formatting.
func (srv *server) GetLoadouts() http.HandlerFunc {
	srv.Log.Info("Getting loadouts")
	return func(w http.ResponseWriter, r *http.Request) {
		loadouts := models.GetStaticLoadouts()
		err := loadouts.ToJSON(w)
		if err != nil {
			srv.Log.Error(errors.Wrap(err, "Unable to marshal JSON into loadout struct"))
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}

/*
ValidateLoadout is a middleware function to determine whether the JSON representation
of type Loadout is valid.

If JSON is valid, ValidateLoadout calls the HandlerFunc it was passed.

If JSON is not valid, ValidateLoadout logs an error and returns http.StatusBadRequest to the caller.
*/
func (srv *server) ValidateLoadout(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loadout := &models.Loadout{}
		err := loadout.FromJSON(r.Body)
		if err != nil {
			srv.Log.Error(errors.Wrap(err, "error deserializing loadout passed as request"))
			http.Error(w, "Error reading product", http.StatusBadRequest)
			return
		}

		err = loadout.Validate()
		if err != nil {
			srv.Log.Error(errors.Wrap(err, "error validating loadout passed as request"))
			http.Error(w, "Error validating product: "+err.Error(), http.StatusBadRequest)
			return
		}
		// add the loadout to the request context for use in next handler
		ctx := context.WithValue(r.Context(), LoadoutKey{}, loadout)
		r = r.WithContext(ctx)
		// call next handler
		h(w, r)
	}
}

//
// import (
// 	"context"
// 	"encoding/json"
// 	"github.com/chancetudor/dubzone-api/internal/logger"
// 	"github.com/chancetudor/dubzone-api/internal/models"
// 	"github.com/gorilla/mux"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"net/http"
// 	"strings"
// 	"time"
// )
//
// /* Loadout endpoints live here
// * functionality for
// 	* Reading all loadouts
// 	* Reading loadouts given a category
// 	* Reading loadouts given a weapon name
// */
//
// func (srv *server) CreateLoadoutEndpoint(response http.ResponseWriter, request *http.Request) {
// 	var loadout models.Loadout
// 	if err := json.NewDecoder(request.Body).Decode(&loadout); err != nil {
// 		srv.respond(response, err, http.StatusInternalServerError)
// 		logger.Error(err, "Error decoding JSON body", "CreateLoadoutEndpoint")
// 		return
// 	}
// 	// put weapon name in ALL CAPS
// 	loadout.Primary.WeaponName = strings.ToUpper(loadout.Primary.WeaponName)
// 	loadout.Secondary.WeaponName = strings.ToUpper(loadout.Secondary.WeaponName)
//
// 	db := srv.Auth.Database
// 	collection := srv.Client.Database(db).Collection(srv.Auth.LoadoutCollection)
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
//
// 	result, err := collection.InsertOne(ctx, loadout)
// 	if err != nil {
// 		srv.respond(response, err, http.StatusInternalServerError)
// 		logger.Error(err, "Error decoding JSON body", "CreateLoadoutEndpoint")
// 		return
// 	}
//
// 	srv.respond(response, result, http.StatusOK)
// }
//
// // ReadLoadoutsEndpoint returns all loadouts
// // GET /loadouts
// func (srv *server) ReadLoadoutsEndpoint(response http.ResponseWriter, request *http.Request) {
// 	query := bson.M{}
// 	loadouts, err := srv.getLoadouts(query)
//
// 	if err != nil {
// 		srv.respond(response, err, http.StatusInternalServerError)
// 		logger.Error(nil, "No loadouts retrieved", "ReadLoadoutsEndpoint")
// 		return
// 	}
//
// 	if loadouts == nil {
// 		srv.respond(response, nil, http.StatusNoContent)
// 		logger.Info("No loadouts found", "ReadLoadoutsByWeaponEndpoint")
// 		return
// 	}
//
// 	srv.respond(response, loadouts, http.StatusOK)
// }
//
// // ReadLoadoutsByMetaEndpoint returns loadouts that are listed as meta in the DB
// // GET /loadouts/meta
// func (srv *server) ReadLoadoutsByMetaEndpoint(response http.ResponseWriter, request *http.Request) {
// 	query := bson.M{"meta_loadout": true}
// 	loadouts, err := srv.getLoadouts(query)
//
// 	if err != nil {
// 		srv.respond(response, err, http.StatusInternalServerError)
// 		logger.Error(err, "No loadouts retrieved", "ReadLoadoutsByCategoryEndpoint")
// 		return
// 	}
//
// 	if loadouts == nil {
// 		srv.respond(response, nil, http.StatusNoContent)
// 		logger.Info("No loadouts found", "ReadLoadoutsByWeaponEndpoint")
// 		return
// 	}
//
// 	srv.respond(response, loadouts, http.StatusOK)
// }
//
// // ReadLoadoutsByCategoryEndpoint returns all loadouts having a primary weapon with a specified category
// // GET /loadouts/category/{cat}
// func (srv *server) ReadLoadoutsByCategoryEndpoint(response http.ResponseWriter, request *http.Request) {
// 	params := mux.Vars(request)
// 	var category string
// 	// if category is smg, we must put it in uppercase
// 	// else, we capitalize the first letter of category
// 	if strings.EqualFold("smg", params["cat"]) {
// 		category = strings.ToUpper(params["cat"])
// 	} else {
// 		category = strings.Title(params["cat"])
// 	}
//
// 	query := bson.M{"primary.category": category}
// 	loadouts, err := srv.getLoadouts(query)
//
// 	if err != nil {
// 		srv.respond(response, err, http.StatusInternalServerError)
// 		logger.Error(err, "No loadouts retrieved", "ReadLoadoutsByCategoryEndpoint")
// 		return
// 	}
//
// 	if loadouts == nil {
// 		srv.respond(response, nil, http.StatusNoContent)
// 		logger.Info("No loadouts found", "ReadLoadoutsByCategoryEndpoint")
// 		return
// 	}
//
// 	srv.respond(response, loadouts, http.StatusOK)
// }
//
// // ReadLoadoutsByWeaponEndpoint returns all loadouts for a specified primary weapon
// // GET /loadouts/weapon/{name}
// func (srv *server) ReadLoadoutsByWeaponEndpoint(response http.ResponseWriter, request *http.Request) {
// 	params := mux.Vars(request)
// 	// weapon_name is stored in the database in ALL CAPS, so we UPPERCASE the query param
// 	weaponName := strings.ToUpper(params["name"])
// 	query := bson.M{"primary.weapon_name": weaponName}
// 	loadouts, err := srv.getLoadouts(query)
//
// 	if err != nil {
// 		srv.respond(response, err, http.StatusInternalServerError)
// 		logger.Error(err, "No loadouts retrieved", "ReadLoadoutsByWeaponEndpoint")
// 		return
// 	}
//
// 	if loadouts == nil {
// 		srv.respond(response, nil, http.StatusNoContent)
// 		logger.Info("No loadouts found", "ReadLoadoutsByWeaponEndpoint")
// 		return
// 	}
//
// 	srv.respond(response, loadouts, http.StatusOK)
// }
//
// // getLoadouts is a helper function to retrieve all loadouts
// // and contains the true logic for querying the database
// func (srv *server) getLoadouts(query bson.M) ([]models.Loadout, error) {
// 	db := srv.Auth.Database
// 	collection := srv.Client.Database(db).Collection(srv.Auth.LoadoutCollection)
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
//
// 	// find all documents using the given bson.M{} query,
// 	// where the bson.D{} query can specify category, weapon, or be empty (find all loadouts)
// 	cursor, err := collection.Find(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer func(cursor *mongo.Cursor, ctx context.Context) {
// 		_ = cursor.Close(ctx)
// 	}(cursor, ctx)
//
// 	var loadouts []models.Loadout
// 	for cursor.Next(ctx) {
// 		var loadout models.Loadout
// 		if err = cursor.Decode(&loadout); err != nil {
// 			return nil, err
// 		}
// 		// append encoded Loadout in []Loadouts
// 		loadouts = append(loadouts, loadout)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}
//
// 	return loadouts, nil
// }
