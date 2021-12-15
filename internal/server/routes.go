package server

import (
	auth "github.com/goji/httpauth"
	"github.com/joho/godotenv"
	"github.com/justinas/alice"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// routes creates a subrouter for each type of model (loadouts and weapons),
// and ties middleware and final handlers to each path.
func (srv *server) routes() {
	srv.log.WithFields(logrus.Fields{"Caller": "routes()", "Message": "Initializing router and adding handlers"}).Info()
	// TODO remove when we have custom auth built
	err := godotenv.Load("/home/chance/GitHub/dubzone-api/dev.env")
	if err != nil {
		srv.log.Error(errors.Wrap(err, "Unable to load dev.env file. Not setting up routes."))
		return
	}
	// loadout specific handlers
	// VERB basePath/loadouts/
	loadoutsRouter := srv.router.PathPrefix("/loadouts").Subrouter().StrictSlash(true)
	loadoutsRouter.Handle("/", alice.
		New(srv.mdl.Log).
		ThenFunc(srv.GetLoadouts())).
		Methods(http.MethodGet).
		Schemes("http", "https").
		Name("GetAllLoadouts")
	// TODO remove the authentication function below in favor of custom auth
	loadoutsRouter.Handle("/", alice.
		New(srv.mdl.Log, auth.SimpleBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD")), srv.mdl.ValidateLoadout).
		ThenFunc(srv.CreateLoadout())).
		Methods(http.MethodGet).
		Schemes("http", "https"). // TODO switch to https in prod
		Name("PostLoadouts")
	loadoutsRouter.Handle("/weapon", alice.
		New(srv.mdl.Log, srv.mdl.ValidateWeaponNameParam).
		ThenFunc(srv.GetLoadoutsByWeapon())).
		Methods(http.MethodGet).
		Queries("name", "{name:[a-zA-Z]+\\d*}"). // match one or more of any a-z char and zero or more of any digit
		Schemes("http", "https").
		Name("GetLoadoutsByWeapon")
	loadoutsRouter.Handle("/weapon", alice.
		New(srv.mdl.Log, srv.mdl.ValidateCategoryParam).
		ThenFunc(srv.GetLoadoutsByCategory())).
		Queries("category", "{category:[a-zA-Z]+[-]?[a-zA-Z]*}"). // match one or more of any a-z char, zero or one hyphen, and zero or more of any a-z char
		Methods(http.MethodGet).
		Schemes("http", "https").
		Name("GetLoadoutsByCategory")
	loadoutsRouter.Handle("/weapon", alice.
		New(srv.mdl.Log, srv.mdl.ValidateGameParam).
		ThenFunc(srv.GetLoadoutsByGame())).
		Queries("game", "{game:[a-zA-Z]+[\\s]?[a-zA-Z]+"). // match one or more of any a-z char, zero or one whitespace, and one or more of any a-z char
		Methods(http.MethodGet).
		Schemes("http", "https").
		Name("GetLoadoutsByGame")
	loadoutsRouter.Handle("/meta", alice.
		New(srv.mdl.Log).
		ThenFunc(srv.GetMetaLoadouts())).
		Methods(http.MethodGet).
		Schemes("http", "https")

	// weapons handlers
	// VERB basePath/weapons/
	weaponsRouter := srv.router.PathPrefix("/weapons").Subrouter().StrictSlash(true)
	weaponsRouter.Handle("/", alice.
		New(srv.mdl.Log).
		ThenFunc(srv.GetWeapons())).
		Methods(http.MethodGet).
		Schemes("http", "https").
		Name("GetAllWeapons")
	weaponsRouter.Handle("/weapon", alice.
		New(srv.mdl.Log, srv.mdl.ValidateWeaponNameParam).
		ThenFunc(srv.GetWeaponsByName())).
		Methods(http.MethodGet).
		Queries("name", "{name:[a-zA-Z]+\\d*}").
		Schemes("http", "https").
		Name("GetWeaponsByName")
	weaponsRouter.Handle("/weapon", alice.
		New(srv.mdl.Log, srv.mdl.ValidateCategoryParam).
		ThenFunc(srv.GetWeaponsByCategory())).
		Queries("category", "{category:[a-zA-Z]+[-]?[a-zA-Z]*}"). // match one or more of any a-z char, zero or one hyphen, and zero or more of any a-z char
		Methods(http.MethodGet).
		Schemes("http", "https").
		Name("GetWeaponsByCategory")
	weaponsRouter.Handle("/weapon", alice.
		New(srv.mdl.Log, srv.mdl.ValidateGameParam).
		ThenFunc(srv.GetWeaponsByGame())).
		Queries("game", "{game:[a-zA-Z]+[\\s]?[a-zA-Z]+}"). // match one or more of any a-z char, zero or one whitespace, and one or more of any a-z char
		Methods(http.MethodGet).
		Schemes("http", "https").
		Name("GetWeaponsByGame")
	weaponsRouter.Handle("/categories", alice.
		New(srv.mdl.Log).
		ThenFunc(srv.GetWeaponCategories())).
		Methods(http.MethodGet).
		Schemes("http", "https")
	weaponsRouter.Handle("/meta", alice.
		New(srv.mdl.Log).
		ThenFunc(srv.GetMetaWeapons())).
		Methods(http.MethodGet).
		Schemes("http", "https")
}
