package server

import (
	"context"
	"github.com/chancetudor/dubzone-api/internal/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
	"time"
)

/* Weapon endpoints live here
* functionality for
	* Reading a single weapon
*/

// ReadWeaponEndpoint returns weapon data for a specified weapon name
// GET /weapon/{weaponname}
// TODO deal with weapon name param containing spaces
func (srv *server) ReadWeaponEndpoint(response http.ResponseWriter, request *http.Request) {
	db := srv.Auth.Database
	collection := srv.Client.Database(db).Collection(srv.Auth.WeaponsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	params := mux.Vars(request)
	weaponName := strings.ToUpper(params["weaponname"])
	var weapon models.Weapon
	// find weapon using given weaponname
	err := collection.FindOne(ctx, bson.D{{"weapon_name", weaponName}}).Decode(&weapon)
	if err != nil {
		srv.respond(response, err, http.StatusInternalServerError)
		// srv.error(err, "Error reading a weapon")
		return
	}

	srv.respond(response, weapon, http.StatusOK)
}

// // ReadWeaponsEndpoint returns all weapons in the Weapons collection
// // GET /weapons
// func (srv *server) ReadWeaponsEndpoint(response http.ResponseWriter, request *http.Request) {
// 	var weapons []models.Weapon
// 	query := bson.M{}
// 	weapons, err := srv.readManyWeapons(query)
//
// 	if err != nil {
// 		srv.respond(response, err, http.StatusInternalServerError)
// 		logger.Error(err, "Error reading weapons", "srv.ReadWeaponsEndpoint()")
// 		return
// 	}
//
// 	srv.respond(response, weapons, http.StatusOK)
// }
//
// // readManyWeapons is srv helper function for ReadWeaponsEndpoint
// // and contains the true logic for querying the database
// // returns a slice of Weapons or an error
// func (srv *server) readManyWeapons(query bson.M) ([]models.Weapon, error) {
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
// 		// srv.error(err, "Error finding weapons in DB")
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
// 			logger.Error(err, "Error decoding weapons cursor into Weapon", "readManyWeapons()")
// 			return nil, err
// 		}
// 		// append encoded Loadout in []Loadouts
// 		weapons = append(weapons, weapon)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		logger.Error(err, "Error with weapons cursor", "readManyWeapons()")
// 		return nil, err
// 	}
//
// 	return weapons, nil
// }
