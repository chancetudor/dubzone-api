package server

import (
	"github.com/justinas/alice"
	"github.com/sirupsen/logrus"
	"net/http"
)

// routes creates a subrouter for each type of request,
// and ties middleware and final handlers to each path.
func (srv *server) routes() {
	srv.log.WithFields(logrus.Fields{"Caller": "routes()", "Message": "Initializing router and adding handlers"}).Info()
	loadoutsRouter := srv.router.PathPrefix("/loadouts/").Subrouter()
	loadoutsRouter.Handle("/", alice.
		New(srv.mdl.Log).
		ThenFunc(srv.GetLoadouts())).
		Schemes("http", "https").
		Methods(http.MethodGet)
	loadoutsRouter.Handle("/", alice.
		New(srv.mdl.Log, srv.mdl.Authenticate, srv.mdl.ValidateLoadout).
		ThenFunc(srv.CreateLoadout())).
		Schemes("http"). // TODO switch to https in prod
		Methods(http.MethodPost)
	loadoutsRouter.Handle("/weapon/{weapon_name}", alice.
		New(srv.mdl.Log, srv.mdl.ValidateNameParam).
		ThenFunc(srv.GetLoadoutsByWeapon())).
		Schemes("http", "https").
		Methods(http.MethodGet)
	loadoutsRouter.Handle("/category/{category}", alice.
		New(srv.mdl.Log, srv.mdl.ValidateCategoryParam).
		ThenFunc(srv.GetLoadoutsByCategory())).
		Schemes("http", "https").
		Methods(http.MethodGet)
	loadoutsRouter.Handle("/meta", alice.
		New(srv.mdl.Log).
		ThenFunc(srv.GetMetaLoadouts())).
		Schemes("http", "https").
		Methods(http.MethodGet)

	// init subrouter just for GET requests and associate middleware (alice) + final handler
	// getRouter := srv.router.Methods(http.MethodGet).Subrouter()

	// loadouts handlers
	// getRouter.Handle("/loadouts", alice.New(srv.mdl.Log).ThenFunc(srv.GetLoadouts())).Schemes("https")
	// getRouter.Handle("/loadouts/category/{category}", alice.
	// 	New(srv.mdl.ValidateCategoryParam).
	// 	ThenFunc(srv.GetLoadoutsByCategory())).
	// 	Schemes("https")
	// getRouter.Handle("/loadouts/weapon/{name}", alice.
	// 	New(srv.mdl.ValidateNameParam).
	// 	ThenFunc(srv.GetLoadoutsByWeapon())).
	// 	Schemes("https")
	// getRouter.Handle("/loadouts/meta", alice.
	// 	New(srv.mdl.Log).
	// 	ThenFunc(srv.GetMetaLoadouts())).
	// 	Schemes("https")

	/*
		weapons handlers
		getRouter.HandleFunc("/weapons/{name}", srv.GetWeaponsByName())
		getRouter.HandleFunc("/weapons/meta", srv.GetMetaWeapons())
		getRouter.HandleFunc("/weapons/{category}", srv.GetWeaponsByCategory())
		getRouter.HandleFunc("/weapons/categories", srv.GetWeaponCategories())
		getRouter.HandleFunc("/weapons/{game}", srv.GetWeaponsByGame())
	*/

	// init subrouter just for POST requests and associate middleware (alice) + final handler
	// postRouter := srv.router.Methods(http.MethodPost).Subrouter()
	// postRouter.Handle("/loadouts", alice.
	// 	New(srv.mdl.Log, srv.mdl.Authenticate, srv.mdl.ValidateLoadout).
	// 	ThenFunc(srv.CreateLoadout())).
	// 	Schemes("https")
}
