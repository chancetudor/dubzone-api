package api

// The router file is used to deal with router handling

import (
	"github.com/gorilla/mux"
	"log"
)

func NewRouter() *mux.Router {
	log.Println("Creating new router: " + "func NewRouter()")
	r := mux.NewRouter().StrictSlash(true)
	//r.Queries()
	//initRouter(r)

	return r
}

// InitRouter creates new router and initializes functions that will handle requests
func InitRouter(r *mux.Router) {
	log.Println("Initializing router: " + "func InitRouter()")
	// TODO FIND A WAY TO ADD PARAMETER QUERIES/FILTERING INSTEAD OF MULTIPLE ENDPOINTS
	// single weapon endpoints -- these return a single weapon, or dmg profile / loadout associated with a weapon
	r.HandleFunc("/weapon/{weaponname}", CreateWeaponEndpoint).Methods("POST")
	r.HandleFunc("/weapon/{weaponname}", ReadWeaponEndpoint).Methods("GET")
	r.HandleFunc("/weapon/{weaponname}", UpdateWeaponEndpoint).Methods("PUT")
	r.HandleFunc("/weapon/{weaponname}", DeleteWeaponEndpoint).Methods("DELETE")
	r.HandleFunc("/dmgprofile/{weaponname}", CreateDamageProfileEndpoint).Methods("POST")
	r.HandleFunc("/dmgprofile/{weaponname}", ReadDamageProfileEndpoint).Methods("GET")
	r.HandleFunc("/dmgprofile/{weaponname}", UpdateDamageProfileEndpoint).Methods("PUT")
	r.HandleFunc("/dmgprofile/{weaponname}", DeleteDamageProfileEndpoint).Methods("DELETE")
	r.HandleFunc("/loadout", CreateLoadoutEndpoint).Methods("POST")
	r.HandleFunc("/loadout/{weaponname}", CreateLoadoutEndpoint).Methods("POST")
	r.HandleFunc("/loadout/{weaponname}", ReadLoadoutEndpoint).Methods("GET")
	r.HandleFunc("/loadout/{weaponname}", UpdateLoadoutEndpoint).Methods("PUT")
	r.HandleFunc("/loadout/{weaponname}", DeleteLoadoutEndpoint).Methods("DELETE")
	// multiple weapon endpoints
	r.HandleFunc("/weapons", ReadWeaponsEndpoint).Methods("GET")
	r.HandleFunc("/dmgprofiles", ReadDamageProfilesEndpoint).Methods("GET")
	r.HandleFunc("/loadouts", ReadLoadoutsEndpoint).Methods("GET")
	r.HandleFunc("/loadouts/{category}", ReadLoadoutsEndpoint).Methods("GET")
}