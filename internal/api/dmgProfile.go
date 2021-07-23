package api

import "net/http"

// DamageProfile endpoints live here TODO FINISH

func CreateDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {

}

func ReadDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {

}

func UpdateDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {

}

func DeleteDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {

}

func ReadDamageProfilesEndpoint(response http.ResponseWriter, request *http.Request) {

}
/*
	// single weapon endpoints
	r.HandleFunc("/dmgprofile", CreateDamageProfileEndpoint).Methods("POST")
	r.HandleFunc("/dmgprofile", ReadDamageProfileEndpoint).Methods("GET")
	r.HandleFunc("/dmgprofile", UpdateDamageProfileEndpoint).Methods("PUT")
	r.HandleFunc("/dmgprofile", DeleteDamageProfileEndpoint).Methods("DELETE")
	// multiple weapon endpoints
	r.HandleFunc("/dmgprofiles", ReadDamageProfileEndpoint).Methods("GET")
 */

// TODO we want to use filtering
// e.g. GET /dmgprofile?weaponID=12345 -- this returns the dmgprofile for specific weapon ID
// OR GET /loadouts?weaponname=XM4 -- this returns the dmgprofile for specific weapon name