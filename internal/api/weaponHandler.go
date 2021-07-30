package api

import "net/http"

// Weapon endpoints live here TODO FINISH

func CreateWeaponEndpoint(response http.ResponseWriter, request *http.Request) {

}

func ReadWeaponEndpoint(response http.ResponseWriter, request *http.Request) {

}

func UpdateWeaponEndpoint(response http.ResponseWriter, request *http.Request) {

}

func DeleteWeaponEndpoint(response http.ResponseWriter, request *http.Request) {

}

func ReadWeaponsEndpoint(response http.ResponseWriter, request *http.Request) {

}

/*
	// single weapon endpoints
	r.HandleFunc("/weapon", CreateWeaponEndpoint).Methods("POST")
	r.HandleFunc("/weapon", ReadWeaponEndpoint).Methods("GET")
	r.HandleFunc("/weapon", UpdateWeaponEndpoint).Methods("PUT")
	r.HandleFunc("/weapon", DeleteWeaponEndpoint).Methods("DELETE")
	// multiple weapon endpoints
	r.HandleFunc("/weapons", ReadWeaponEndpoint).Methods("GET")
*/

// a weapon has associated loadouts and associated damage profiles, so when a user is retrieving a weapon, we want to
// retrieve the associated data too
// TODO we want to use filtering
// e.g. GET /loadouts?weaponID=12345 -- this returns the loadouts for specific weapon ID
// OR GET /loadouts?weaponname=XM4
