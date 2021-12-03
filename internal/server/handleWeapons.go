package server

//
// import (
// 	"context"
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
// /* Weapon endpoints live here
// * functionality for
// 	* Reading a single weapon
// 	* Reading a list of meta weapons
// 	* Reading a list of weapons based on category
// */
//
// // ReadWeaponEndpoint returns weapon data for a specified weapon name
// // GET /weapon/{name}
// func (srv *server) ReadWeaponEndpoint(response http.ResponseWriter, request *http.Request) {
// 	db := srv.Auth.Database
// 	collection := srv.Client.Database(db).Collection(srv.Auth.WeaponsCollection)
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
//
// 	params := mux.Vars(request)
// 	weaponName := strings.ToUpper(params["name"])
// 	var weapon models.Weapon
// 	// find weapon using given weaponname
// 	err := collection.FindOne(ctx, bson.D{{"weapon_name", weaponName}}).Decode(&weapon)
// 	if err != nil {
// 		srv.respond(response, err, http.StatusInternalServerError)
// 		logger.Error(err, "No weapons retrieved", "ReadWeaponEndpoint")
// 		return
// 	}
//
// 	srv.respond(response, weapon, http.StatusOK)
// }
//
// // ReadWeaponsByMetaEndpoint returns a list of weapons that are considered meta
// // GET /weapons/meta
// func (srv *server) ReadWeaponsByMetaEndpoint(response http.ResponseWriter, request *http.Request) {
// 	query := bson.M{"meta_weapon": true}
// 	weapons, err := srv.getWeapons(query)
//
// 	if err != nil {
// 		srv.respond(response, err, http.StatusInternalServerError)
// 		logger.Error(err, "Error reading weapons", "ReadWeaponsByMetaEndpoint")
// 		return
// 	}
//
// 	if weapons == nil {
// 		srv.respond(response, nil, http.StatusNoContent)
// 		logger.Info("No weapons found", "ReadWeaponsByMetaEndpoint")
// 		return
// 	}
//
// 	srv.respond(response, weapons, http.StatusOK)
// }
//
// // ReadWeaponsByCategoryEndpoint returns a list of weapons that are in a specified category
// // GET /weapons/{cat}
// func (srv *server) ReadWeaponsByCategoryEndpoint(response http.ResponseWriter, request *http.Request) {
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
// 	query := bson.M{"category": category}
// 	weapons, err := srv.getWeapons(query)
//
// 	if err != nil {
// 		srv.respond(response, err, http.StatusInternalServerError)
// 		logger.Error(err, "Error reading weapons", "ReadWeaponsByCategoryEndpoint")
// 		return
// 	}
//
// 	if weapons == nil {
// 		srv.respond(response, nil, http.StatusNoContent)
// 		logger.Info("No weapons found", "ReadWeaponsByCategoryEndpoint")
// 		return
// 	}
//
// 	srv.respond(response, weapons, http.StatusOK)
// }
//
// // GetWeaponCategories returns a list of weapon categories
// func (srv *server) GetWeaponCategories(response http.ResponseWriter, request *http.Request) {
// 	categories := []string{"SMG", "Range", "Sniper", "Sniper Support", "Fully Loaded"}
// 	srv.respond(response, categories, http.StatusOK)
// }
//
// // getWeapons is a helper function for ReadWeaponsEndpoint
// // and contains the true logic for querying the database
// // returns a slice of Weapons or an error
// func (srv *server) getWeapons(query bson.M) ([]models.Weapon, error) {
// 	db := srv.Auth.Database
// 	collection := srv.Client.Database(db).Collection(srv.Auth.WeaponsCollection)
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
//
// 	// find all documents using the given bson.M{} query,
// 	// where the bson.M{} query can specify the game srv weapon is from
// 	// or the bson.M{} query can be empty (find all weapons)
// 	cursor, err := collection.Find(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer func(cursor *mongo.Cursor, ctx context.Context) {
// 		_ = cursor.Close(ctx)
// 	}(cursor, ctx)
//
// 	var weapons []models.Weapon
// 	// iterate through cursor and encode documents into Loadout struct
// 	for cursor.Next(ctx) {
// 		var weapon models.Weapon
// 		if err = cursor.Decode(&weapon); err != nil {
// 			return nil, err
// 		}
// 		// append encoded Loadout to []Loadouts
// 		weapons = append(weapons, weapon)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}
//
// 	return weapons, nil
// }
