package server

import (
	"context"
	"github.com/chancetudor/dubzone-api/internal/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strings"
	"time"
)

/* Damage profile endpoints live here
* functionality for
	* Reading a given weapon's damage profile (all ranges, 2 ranges, or one range)
	* Updating a given weapon's damage profile (all ranges, 2 ranges, or one range)
*/

// ReadCloseDamageProfile
// GET /dmgprofile/{weaponname}/close
func (srv *server) ReadCloseDamageProfile(response http.ResponseWriter, request *http.Request) {

}

// GET /dmgprofile/{weaponname}/mid
func (srv *server) ReadMidDamageProfile(response http.ResponseWriter, request *http.Request) {

}

// GET /dmgprofile/{weaponname}/far
func (srv *server) ReadFarDamageProfile(response http.ResponseWriter, request *http.Request) {

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
	var dmgProfile models.DamageProfile
	projection := bson.D{
		{"_id", 0},
		{"game_from", 0},
		{"rpm", 0},
		{"bullet_velocity", 0},
	}

	err := collection.FindOne(ctx, bson.D{
		{"weapon_name", weaponName},
	},
		options.FindOne().SetProjection(projection)).Decode(&dmgProfile)
	if err != nil {
		srv.respond(response, nil, http.StatusInternalServerError)
		srv.error(err, " Error finding weapon in database", "ReadDamageProfiles")
		return
	}

	srv.respond(response, dmgProfile, http.StatusOK)
}

// func (srv *server) CreateDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {
//
// }

// func (srv *server) DeleteDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {
//
// }
