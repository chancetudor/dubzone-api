package server

import (
	"context"
	"github.com/chancetudor/dubzone-api/internal/logger"
	"github.com/chancetudor/dubzone-api/internal/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
	"time"
)

/* Damage profile endpoints live here
* functionality for
	* Reading a given weapon's damage profile (all ranges, 2 ranges, or one range)
	* Updating a given weapon's damage profile (all ranges, 2 ranges, or one range)
*/

// ReadCloseDamageProfile returns only the close range damage profile for weapon requested
// GET /dmgprofile/{weaponname}/close
func (srv *server) ReadCloseDamageProfile(response http.ResponseWriter, request *http.Request) {
	db := srv.Auth.Database
	collection := srv.Client.Database(db).Collection(srv.Auth.WeaponsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	params := mux.Vars(request)
	weaponName := strings.ToUpper(params["weaponname"])
	var weapon models.Weapon

	err := collection.FindOne(ctx, bson.D{{"weapon_name", weaponName}}).Decode(&weapon)

	if err != nil {
		srv.respond(response, err, http.StatusInternalServerError)
		logger.Error(err, "Error finding weapon in database", "ReadDamageProfiles")
		return
	}

	srv.respond(response, weapon.DamageProfile.CloseRange, http.StatusOK)
}

// ReadMidDamageProfile returns only the mid range damage profile for specified weapon
// GET /dmgprofile/{weaponname}/mid
func (srv *server) ReadMidDamageProfile(response http.ResponseWriter, request *http.Request) {
	db := srv.Auth.Database
	collection := srv.Client.Database(db).Collection(srv.Auth.WeaponsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	params := mux.Vars(request)
	weaponName := strings.ToUpper(params["weaponname"])
	var weapon models.Weapon

	err := collection.FindOne(ctx, bson.D{{"weapon_name", weaponName}}).Decode(&weapon)

	if err != nil {
		srv.respond(response, err, http.StatusInternalServerError)
		logger.Error(err, "Error finding weapon in database", "ReadDamageProfiles")
		return
	}

	srv.respond(response, weapon.DamageProfile.MidRange, http.StatusOK)
}

// ReadFarDamageProfile returns only the far range damage profile for specified weapon
// GET /dmgprofile/{weaponname}/far
func (srv *server) ReadFarDamageProfile(response http.ResponseWriter, request *http.Request) {
	db := srv.Auth.Database
	collection := srv.Client.Database(db).Collection(srv.Auth.WeaponsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	params := mux.Vars(request)
	weaponName := strings.ToUpper(params["weaponname"])
	var weapon models.Weapon

	err := collection.FindOne(ctx, bson.D{{"weapon_name", weaponName}}).Decode(&weapon)

	if err != nil {
		srv.respond(response, err, http.StatusInternalServerError)
		logger.Error(err, "Error finding weapon in database", "ReadDamageProfiles")
		return
	}

	srv.respond(response, weapon.DamageProfile.FarRange, http.StatusOK)
}

// PUT /dmgprofile/{weaponname}
func (srv *server) UpdateDamageProfile(response http.ResponseWriter, request *http.Request) {

}

// PUT /dmgprofile/{weaponname}/close
func (srv *server) UpdateCloseDamageProfile(response http.ResponseWriter, request *http.Request) {

}

// PUT /dmgprofile/{weaponname}/mid
func (srv *server) UpdateMidDamageProfile(response http.ResponseWriter, request *http.Request) {

}

// PUT /dmgprofile/{weaponname}/far
func (srv *server) UpdateFarDamageProfile(response http.ResponseWriter, request *http.Request) {

}

// ReadDamageProfiles takes a weapon name as a parameter and returns
// all damage profiles (close, mid, far) for that weapon
// GET /dmgprofile/{weaponname}
func (srv *server) ReadDamageProfiles(response http.ResponseWriter, request *http.Request) {
	db := srv.Auth.Database
	collection := srv.Client.Database(db).Collection(srv.Auth.WeaponsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	params := mux.Vars(request)
	weaponName := strings.ToUpper(params["weaponname"])
	var weapon models.Weapon

	err := collection.FindOne(ctx, bson.D{{"weapon_name", weaponName}}).Decode(&weapon)

	if err != nil {
		srv.respond(response, err, http.StatusInternalServerError)
		logger.Error(err, "Error finding weapon in database", "ReadDamageProfiles")
		return
	}

	srv.respond(response, weapon.DamageProfile, http.StatusOK)
}

// func (srv *server) CreateDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {
//
// }

// func (srv *server) DeleteDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {
//
// }
