package server

import (
	"context"
	"encoding/json"
	"github.com/chancetudor/dubzone-api/internal/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strings"
	"time"
)

/* Weapon endpoints live here
* functionality for
	* Creating srv single weapon
	* Reading srv single weapon
	* Reading multiple weapons, given game name, or returning all weapons
*/

// CreateWeaponEndpoint creates srv single new weapon in the Weapons collection
// POST /weapon
func (srv *server) CreateWeaponEndpoint(response http.ResponseWriter, request *http.Request) {
	var weapon models.Weapon
	err := json.NewDecoder(request.Body).Decode(&weapon)
	weapon.WeaponName = strings.ToUpper(weapon.WeaponName)

	db := srv.Auth.Database
	collection := srv.Client.Database(db).Collection(srv.Auth.WeaponsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer client.Disconnect(ctx)
	defer cancel()

	_, err = collection.InsertOne(ctx, weapon)
	if err != nil {
		srv.respond(response, nil, http.StatusInternalServerError)
		srv.error(err, " Inserting weapon into collection", " CreateWeaponEndpoint")
		return
	}

	srv.respond(response, weapon.WeaponName, http.StatusOK)
}

// ReadWeaponEndpoint returns weapon data for srv specified weapon name
// GET /weapon/{weaponname}
func (srv *server) ReadWeaponEndpoint(response http.ResponseWriter, request *http.Request) {
	db := srv.Auth.Database
	collection := srv.Client.Database(db).Collection(srv.Auth.WeaponsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	params := mux.Vars(request)
	weaponName := strings.ToUpper(params["weaponname"])
	var weapon models.Weapon
	projection := bson.D{
		{"_id", 0},
	}
	// find weapon using given weaponname
	err := collection.FindOne(ctx, bson.D{
		{"weapon_name", weaponName},
	},
		options.FindOne().SetProjection(projection)).Decode(&weapon)
	if err != nil {
		srv.respond(response, nil, http.StatusInternalServerError)
		srv.error(err, " Finding weapon in database", " ReadWeaponEndpoint")
		return
	}

	srv.respond(response, weapon, http.StatusOK)
}

// ReadWeaponsEndpoint returns all weapons in the Weapons collection
// GET /weapons
func (srv *server) ReadWeaponsEndpoint(response http.ResponseWriter, request *http.Request) {
	var weapons []models.Weapon
	query := bson.M{}
	weapons = srv.readManyWeapons(query)

	if weapons == nil {
		srv.respond(response, weapons, http.StatusInternalServerError)
		srv.error(nil, " No weapons retrieved", " ReadWeaponsEndpoint")
		return
	}

	srv.respond(response, weapons, http.StatusOK)
}

// ReadWeaponsByGameEndpoint returns all weapons from specified game
// GET /weapons/{game}
func (srv *server) ReadWeaponsByGameEndpoint(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	game := strings.ToUpper(params["game"])
	var weapons []models.Weapon
	query := bson.M{"game_from": game}
	weapons = srv.readManyWeapons(query)

	if weapons == nil {
		srv.respond(response, weapons, http.StatusInternalServerError)
		srv.error(nil, " No weapons retrieved", " ReadWeaponsByGameEndpoint")
		return
	}

	srv.respond(response, weapons, http.StatusOK)
}

// readManyWeapons is srv helper function for ReadWeaponsEndpoint
// and contains the true logic for querying the database
func (srv *server) readManyWeapons(query bson.M) []models.Weapon {
	db := srv.Auth.Database
	collection := srv.Client.Database(db).Collection(srv.Auth.WeaponsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// find all documents using the given bson.M{} query,
	// where the bson.M{} query can specify the game srv weapon is from
	// or the bson.M{} query can be empty (find all weapons)
	cursor, err := collection.Find(ctx, query)
	if err != nil {
		srv.error(err, " Finding weapons in database", " readManyWeapons")
		return nil
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		_ = cursor.Close(ctx)
	}(cursor, ctx)

	var weapons []models.Weapon
	// TODO optimize for loop -- maybe pagination?
	// iterate through cursor and encode documents into Loadout struct
	for cursor.Next(ctx) {
		var weapon models.Weapon
		if err = cursor.Decode(&weapon); err != nil {
			srv.error(err, " Decoding cursor into Weapon", " readManyWeapons")
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
	* we probably don't need to delete srv weapon
func DeleteWeaponEndpoint(response http.ResponseWriter, request *http.Request) {

}
*/

/*
	* we probably don't need to update srv weapon
// UpdateWeaponEndpoint takes srv specified weapon and srv series of parameters
// and updates that weapon's parameters to the given ones
func UpdateWeaponEndpoint(response http.ResponseWriter, request *http.Request) {

}
*/
