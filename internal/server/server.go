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
	log    *logrus.Logger
	mdl    *middleware.Middleware
}

// NewServer returns a pointer to a type server that has a mux router
// configured with endpoints, an in-memory cache
// set to 5 minute expiration time and 10 minute clean-up time,
// and a pointer to a logger that's passed in.
//
// The server contains a pointer to type middleware.Middleware,
// which is passed the same logs as the server.
// This allows middleware, http.HandlerFuncs,
// to logs any errors to the same logs file as the server.
func NewServer(l *logrus.Logger) *server {
	api := &server{
		router: newRouter(),
		cache:  cache.New(5*time.Minute, 10*time.Minute),
		log:    l,
		mdl:    middleware.NewMiddleware(l),
	}
	api.routes()

	return api
}

// Start simply runs http.ListenAndServe on the passed-in bind address.
func (srv *server) Start(port string) {
	srv.log.Info("Starting server on port " + port)
	srv.log.Fatal(http.ListenAndServe(port, srv.router))
}

// newRouter creates a new Gorilla mux with appropriate options
func newRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	return r
}
