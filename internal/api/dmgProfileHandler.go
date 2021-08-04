package api

import (
	"context"
	"encoding/json"
	"github.com/chancetudor/dubzone-api/internal/models"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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

// ReadDamageProfileEndpoint takes a required weapon name to return its damage profiles
func (a *API) ReadDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	db := a.Auth.Database
	collection := a.Client.Database(db).Collection(a.Auth.WeaponsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	projection := bson.D{
		{"damage_profile", 1},
		{"_id", 0},
	}
	params := mux.Vars(request)
	weaponName := strings.ToUpper(params["weaponname"])
	query := bson.D{{"weapon_name", weaponName}}
	var dmgProfiles models.DamageProfile

	err := collection.FindOne(ctx, query, options.FindOne().SetProjection(projection)).Decode(&dmgProfiles)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		log.WithFields(log.Fields{
			"func":  "ReadWeaponEndpoint()",
			"event": "Finding weapon in database",
		}).Error(err)
		return
	}

	err = json.NewEncoder(response).Encode(dmgProfiles)
}

// ReadCloseDamageProfileEndpoint takes a required weapon name to return its close-range damage profile
func (a *API) ReadCloseDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {

}

// ReadMidDamageProfileEndpoint takes a required weapon name to return its mid-range damage profile
func (a *API) ReadMidDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {

}

// ReadFarDamageProfileEndpoint takes a required weapon name to return its far-range damage profile
func (a *API) ReadFarDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {

}

// UpdateCloseDamageProfileEndpoint takes a required weapon name to update its close-range damage profile
func (a *API) UpdateCloseDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {

}

// UpdateMidDamageProfileEndpoint takes a required weapon name to update its mid-range damage profile
func (a *API) UpdateMidDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {

}

// UpdateFarDamageProfileEndpoint takes a required weapon name to update its far-range damage profile
func (a *API) UpdateFarDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {

}

// func (a *API) ReadDamageProfilesEndpoint(response http.ResponseWriter, request *http.Request) {
//
// }

// func (a *API) CreateDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {
//
// }

// func (a *API) DeleteDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {
//
// }

