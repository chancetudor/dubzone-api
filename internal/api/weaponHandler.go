package api

import "net/http"

const CollectionName = "Weapons"

/* Loadout endpoints live here
* functionality for
	* Creating a single weapon
	* Reading a single weapon
	* Updating a single weapon
	* Reading multiple weapons, given weapon name, or returning all weapons
*/

/*
	// single weapon endpoints
	r.HandleFunc("/weapon", CreateWeaponEndpoint).Methods("POST")
	r.HandleFunc("/weapon", ReadWeaponEndpoint).Methods("GET")
	r.HandleFunc("/weapon", UpdateWeaponEndpoint).Methods("PUT")
	// multiple weapon endpoints
	r.HandleFunc("/weapons", ReadWeaponEndpoint).Methods("GET")
*/

func CreateWeaponEndpoint(response http.ResponseWriter, request *http.Request) {

}

func ReadWeaponEndpoint(response http.ResponseWriter, request *http.Request) {

}

func UpdateWeaponEndpoint(response http.ResponseWriter, request *http.Request) {

}

func ReadWeaponsEndpoint(response http.ResponseWriter, request *http.Request) {

}

/*
	* we probably don't need to delete a weapon
func DeleteWeaponEndpoint(response http.ResponseWriter, request *http.Request) {

}
*/

