package server

import (
	"net/http"
)

/* Damage profile endpoints live here
* functionality for
	* Reading a given weapon's damage profile (all ranges, 2 ranges, or one range)
	* Updating a given weapon's damage profile (all ranges, 2 ranges, or one range)
*/

func (srv *server) ReadDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {

}

func (srv *server) UpdateDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {

}

func (srv *server) ReadDamageProfilesEndpoint(response http.ResponseWriter, request *http.Request) {

}

// func (srv *server) CreateDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {
//
// }

// func (srv *server) DeleteDamageProfileEndpoint(response http.ResponseWriter, request *http.Request) {
//
// }
