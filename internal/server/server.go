// Package classification Dubzone API
//
// A service that communicates meta weapons and loadouts, stored in a database, to a consumer in JSON format.
//
//     Schemes: https
//     Host: localhost
//     BasePath: /
//     Version: 0.0.2
//     Contact: Chance Tudor<hi@cmtudor.me> https://cmtudor.me
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package server

import (
	"github.com/chancetudor/dubzone-api/middleware"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type server struct {
	router *mux.Router
	cache  *cache.Cache
	Log    *logrus.Logger
	mdl    *middleware.Middleware
}

// NewServer returns a pointer to a type server that has a mux router
// configured with endpoints, an in-memory cache
// set to 5 minute expiration time and 10 minute clean-up time,
// and a pointer to a logger that's passed in.
//
// The server contains a pointer to type middleware.Middleware,
// which is passed the same log as the server.
// This allows middleware, basically http.HandlerFuncs,
// to log any errors to the same log file as the server.
func NewServer(l *logrus.Logger) *server {
	api := &server{
		router: newRouter(),
		cache:  cache.New(5*time.Minute, 10*time.Minute),
		Log:    l,
		mdl:    middleware.NewMiddleware(l),
	}
	api.routes()

	return api
}

// Start simply runs http.ListenAndServe on the passed-in bind address.
func (srv *server) Start(port string) error {
	srv.Log.Info("Starting server on port " + port)
	return http.ListenAndServe(port, srv.router)
}

// newRouter creates a new Gorilla mux with appropriate options
func newRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	return r
}

// routes creates a subrouter for each type of request,
// adds middleware validation, and ties handlers to each path.
func (srv *server) routes() {
	srv.Log.Info("Initializing router & adding handlers", "InitRouter()")
	// init subrouter just for GET requests and associate handlers
	getRouter := srv.router.Methods(http.MethodGet).Subrouter()

	// loadouts handlers
	getRouter.HandleFunc("/loadouts", srv.GetLoadouts())
	// getRouter.HandleFunc("/loadouts/category/{cat}", srv.GetLoadoutsByCategory())
	// getRouter.HandleFunc("/loadouts/weapon/{name}", srv.GetLoadoutsByWeapon())
	getRouter.HandleFunc("/loadouts/meta", srv.GetMetaLoadouts())

	// weapons handlers
	// getRouter.HandleFunc("/weapons/{name}", srv.GetWeaponsByName())
	// getRouter.HandleFunc("/weapons/meta", srv.GetMetaWeapons())
	// getRouter.HandleFunc("/weapons/{cat}", srv.GetWeaponsByCategory())
	// getRouter.HandleFunc("/weapons/categories", srv.GetWeaponCategories())

	// init subrouter just for POST requests and associate handlers
	postRouter := srv.router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/loadouts", srv.mdl.ValidateLoadout(srv.CreateLoadout()))
}
