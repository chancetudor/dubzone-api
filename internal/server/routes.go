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
	err := godotenv.Load("/Users/chancetudor/GitHub/dubzone-api/dev.env")
	if err != nil {
		srv.log.Error(errors.Wrap(err, "Unable to load dev.env file"))
		return
	}
	loadoutsRouter := srv.router.PathPrefix("/loadouts").Subrouter()

	loadoutsRouter.Handle("/", alice.
		New(srv.mdl.Log).
		ThenFunc(srv.GetLoadouts())).
		Schemes("http", "https").
		Methods(http.MethodGet)
	// TODO remove the authentication function below in favor of custom auth
	loadoutsRouter.Handle("/", alice.
		New(srv.mdl.Log, auth.SimpleBasicAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD")), srv.mdl.ValidateLoadout).
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

	// weapons handlers
	weaponsRouter := srv.router.PathPrefix("/weapons").Subrouter()
	weaponsRouter.Handle("/", alice.
		New(srv.mdl.Log).
		ThenFunc(srv.GetWeapons())).
		Schemes("http", "https").
		Methods(http.MethodGet)
	weaponsRouter.Handle("/weapon/{weapon_name}", alice.
		New(srv.mdl.Log, srv.mdl.ValidateNameParam).
		ThenFunc(srv.GetWeaponsByName())).
		Schemes("http", "https").
		Methods(http.MethodGet)
	weaponsRouter.Handle("/meta", alice.
		New(srv.mdl.Log).
		ThenFunc(srv.GetMetaWeapons())).
		Schemes("http", "https").
		Methods(http.MethodGet)
	weaponsRouter.Handle("/category/{category}", alice.
		New(srv.mdl.Log, srv.mdl.ValidateCategoryParam).
		ThenFunc(srv.GetWeaponsByCategory())).
		Schemes("http", "https").
		Methods(http.MethodGet)
	weaponsRouter.Handle("/categories", alice.
		New(srv.mdl.Log).
		ThenFunc(srv.GetWeaponCategories())).
		Schemes("http", "https").
		Methods(http.MethodGet)
	weaponsRouter.Handle("/game/{game}", alice.
		New(srv.mdl.Log, srv.mdl.ValidateGameParam).
		ThenFunc(srv.GetWeaponsByGame())).
		Schemes("http", "https").
		Methods(http.MethodGet)
}
