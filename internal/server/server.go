package server

import (
	"context"
	"encoding/json"
	"github.com/chancetudor/dubzone-api/internal/auth"
	"github.com/chancetudor/dubzone-api/internal/logger"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

type server struct {
	Router *mux.Router
	Client *mongo.Client
	Auth   *auth.MongoAuth
}

func NewServer() *server {
	api := &server{
		Router: newRouter(),
		Client: newClient(),
		Auth:   auth.NewAuth(),
	}
	api.initRouter()

	return api
}

// newRouter creates a new Gorilla mux with appropriate options
func newRouter() *mux.Router {
	logger.Info("Creating new router", "NewRouter()")
	r := mux.NewRouter().StrictSlash(true)

	return r
}

// newClient creates a new mongo client with appropriate authentication
func newClient() *mongo.Client {
	logger.Info("Creating new client", "NewClient()")

	mongoAuth := auth.NewAuth()
	clientOptions := options.Client().
		ApplyURI(mongoAuth.URI)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Fatal(err, "Connecting to mongoDB", "NewClient()")
	}

	return client
}

// respond is a helper function to abstract HTTP responses
func (srv *server) respond(response http.ResponseWriter, data interface{}, status int) {
	response.Header().Add("content-type", "application/json")
	// switch on the type of interface{} passed in
	// if data is type error, respond with an error
	// otherwise, encode response
	switch d := data.(type) {
	case error:
		http.Error(response, d.Error(), status)
		return
	default:
		response.WriteHeader(status)
		_ = json.NewEncoder(response).Encode(data)
	}
}

// DisconnectClient is a helper function to disconnect mongo client
func (srv *server) DisconnectClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			logger.Error(err, "Client disconnect", "DisconnectClient()")
		}
	}(srv.Client, ctx)

	logger.Info("DISCONNECTING CLIENT", "DisconnectClient()")
}

// initRouter ties handlers to router
func (srv *server) initRouter() {
	logger.Info("Initializing router & adding handlers", "InitRouter()")
	srv.Router.HandleFunc("/loadouts", srv.ReadLoadoutsEndpoint).Methods("GET")
	srv.Router.HandleFunc("/loadout", srv.CreateLoadoutEndpoint).Methods("POST")
	srv.Router.HandleFunc("/loadouts/category/{cat}", srv.ReadLoadoutsByCategoryEndpoint).Methods("GET")
	srv.Router.HandleFunc("/loadouts/weapon/{name}", srv.ReadLoadoutsByWeaponEndpoint).Methods("GET")
	srv.Router.HandleFunc("/loadouts/meta", srv.ReadLoadoutsByMetaEndpoint).Methods("GET")

	srv.Router.HandleFunc("/weapon/{name}", srv.ReadWeaponEndpoint).Methods("GET")
	srv.Router.HandleFunc("/weapons/meta", srv.ReadWeaponsByMetaEndpoint).Methods("GET")
	srv.Router.HandleFunc("/weapons/{cat}", srv.ReadWeaponsByCategoryEndpoint).Methods("GET")
	srv.Router.HandleFunc("/weapons/categories", srv.GetWeaponCategories).Methods("GET")
}