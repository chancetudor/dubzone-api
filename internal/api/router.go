package api

// The router file is used to deal with router handling

import (
	"github.com/gorilla/mux"
	"log"
)

func NewRouter() *mux.Router {
	log.Println("Creating new router: " + "func NewRouter()")
	r := mux.NewRouter()
	InitRouter(r)

	return r
}

// InitRouter creates new router and initializes functions that will handle requests
func InitRouter(r *mux.Router) {
	log.Println("Initializing router: " + "func InitRouter()")
	/* handle functions for
	* creating weapon
	* reading weapon
	* updating weapon
	* deleting weapon
	* creating damage profile
	* reading damage profile
	* updating damage profile
	* deleting damage profile
	* creating loadout
	* reading loadout
	* updating loadout
	* deleting loadout
	 */
	/*
	r.HandleFunc("/weapon/{weaponname}", CreateWeaponEndpoint).Methods("POST")
	r.HandleFunc("/weapon/{weaponname}", ReadWeaponEndpoint).Methods("GET")
	r.HandleFunc("/weapon/{weaponname}", UpdateWeaponEndpoint).Methods("PUT")
	r.HandleFunc("/weapon/{weaponname}", DeleteWeaponEndpoint).Methods("DELETE")
	r.HandleFunc("/dmgprofile/{weaponname}", CreateDamageProfileEndpoint).Methods("POST")
	r.HandleFunc("/dmgprofile/{weaponname}", ReadDamageProfileEndpoint).Methods("GET")
	r.HandleFunc("/dmgprofile/{weaponname}", UpdateDamageProfileEndpoint).Methods("PUT")
	r.HandleFunc("/dmgprofile/{weaponname}", DeleteDamageProfileEndpoint).Methods("DELETE")
	r.HandleFunc("/loadout", CreateLoadoutEndpoint).Methods("POST")
	r.HandleFunc("/loadout", ReadLoadoutEndpoint).Methods("GET")
	r.HandleFunc("/loadout", UpdateLoadoutEndpoint).Methods("PUT")
	r.HandleFunc("/loadout", DeleteLoadoutEndpoint).Methods("DELETE")
	*/
}