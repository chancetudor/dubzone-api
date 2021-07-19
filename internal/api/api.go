package api

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// The API file ties together the router and client

type API struct {
	Router *mux.Router
	Client *mongo.Client
}

func New() *API {
	log.Println("Creating new API struct: " + "func api.New()")
	return &API {
		Router: NewRouter(),
		Client: NewClient(),
	}
}