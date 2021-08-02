package api

// The router file is used to deal with router handling

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func NewRouter() *mux.Router {
	log.Println("Creating new router: " + "func NewRouter()")
	r := mux.NewRouter().StrictSlash(true)

	return r
}

// InitRouter creates new router and initializes functions that will handle requests
func InitRouter(r *mux.Router) {
	log.Println("Initializing router, adding handlers: " + "func InitRouter()")
	// single weapon endpoints, which deal with a single weapon
	r.HandleFunc("/weapon", CreateWeaponEndpoint).Methods("POST")
	r.HandleFunc("/weapon/{weaponname}", ReadWeaponEndpoint).Methods("GET")
	// single dmgProfile endpoints, which deal with a single dmgProfile for a given weapon
	r.HandleFunc("/dmgprofile/{weaponname}", CreateDamageProfileEndpoint).Methods("POST")
	r.HandleFunc("/dmgprofile/{weaponname}", ReadDamageProfileEndpoint).Methods("GET")
	r.HandleFunc("/dmgprofile/{weaponname}", UpdateDamageProfileEndpoint).Methods("PUT")
	r.HandleFunc("/dmgprofile/{weaponname}", DeleteDamageProfileEndpoint).Methods("DELETE")
	// single loadout endpoints, which deal with a single loadout
	r.HandleFunc("/loadout", CreateLoadoutEndpoint).Methods("POST")
	// returns multiple weapons
	r.HandleFunc("/weapons", ReadWeaponsEndpoint).Methods("GET")
	// returns multiple dmgProfiles
	r.HandleFunc("/dmgprofiles", ReadDamageProfilesEndpoint).Methods("GET")
	// returns multiple loadouts
	r.HandleFunc("/loadouts", ReadLoadoutsEndpoint).Methods("GET")
}
