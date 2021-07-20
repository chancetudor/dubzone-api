package api

// DamageProfile endpoints live here
/*
r.HandleFunc("/dmgprofile", CreateDamageProfileEndpoint).Methods("POST")
	r.HandleFunc("/dmgprofile", ReadDamageProfileEndpoint).Methods("GET")
	r.HandleFunc("/dmgprofile", UpdateDamageProfileEndpoint).Methods("PUT")
	r.HandleFunc("/dmgprofile", DeleteDamageProfileEndpoint).Methods("DELETE")
 */

// TODO we want to use filtering
// e.g. GET /dmgprofile?weaponID=12345 -- this returns the dmgprofile for specific weapon ID
// OR GET /loadouts?weaponname=XM4 -- this returns the dmgprofile for specific weapon name