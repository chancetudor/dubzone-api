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
	logger.Debug("Creating new router", "NewRouter()")
	r := mux.NewRouter().StrictSlash(true) //.UseEncodedPath() TODO add in and unescape paramters where necesse est

	return r
}

// newClient creates a new mongo client with appropriate authentication
func newClient() *mongo.Client {
	logger.Debug("Creating new client", "NewClient()")

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
		if err := json.NewEncoder(response).Encode(data); err != nil {
			http.Error(response, err.Error(), status)
			logger.Error(err, "Encoding data into JSON response", "srv.response()")
		}
	}
}

func (srv *server) DisconnectClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			logger.Error(err, "Client disconnect", "DisconnectClient()")
		}
	}(srv.Client, ctx)

	logger.Debug("DISCONNECTING CLIENT", "DisconnectClient()")
}

// initRouter initializes handler funcs on router
func (srv *server) initRouter() {
	logger.Debug("Initializing router, adding handlers", "InitRouter()")
	// returns multiple loadouts
	srv.Router.HandleFunc("/loadouts", srv.ReadLoadoutsEndpoint).Methods("GET")
	srv.Router.HandleFunc("/loadouts/{category}", srv.ReadLoadoutsEndpoint).Methods("GET")
	srv.Router.HandleFunc("/loadouts/{weaponname}", srv.ReadLoadoutsEndpoint).Methods("GET")

	srv.Router.HandleFunc("/weapon/{weaponname}", srv.ReadWeaponEndpoint).Methods("GET")
	// single weapon endpoints, which deal with a single weapon
	// srv.Router.HandleFunc("/weapon", srv.CreateWeaponEndpoint).Methods("POST")
	//
	// // single dmgProfile endpoints, which deal with a single dmgProfile for a given weapon
	// srv.Router.HandleFunc("/dmgprofile/{weaponname}/close", srv.ReadCloseDamageProfile).Methods("GET")
	// srv.Router.HandleFunc("/dmgprofile/{weaponname}/mid", srv.ReadMidDamageProfile).Methods("GET")
	// srv.Router.HandleFunc("/dmgprofile/{weaponname}/far", srv.ReadFarDamageProfile).Methods("GET")
	// srv.Router.HandleFunc("/dmgprofile/{weaponname}", srv.UpdateDamageProfile).Methods("PUT")
	// srv.Router.HandleFunc("/dmgprofile/{weaponname}/close", srv.UpdateCloseDamageProfile).Methods("PUT")
	// srv.Router.HandleFunc("/dmgprofile/{weaponname}/mid", srv.UpdateMidDamageProfile).Methods("PUT")
	// srv.Router.HandleFunc("/dmgprofile/{weaponname}/far", srv.UpdateFarDamageProfile).Methods("PUT")
	// single loadout endpoints, which deal with a single loadout
	// returns multiple weapons
	// srv.Router.HandleFunc("/weapons", srv.ReadWeaponsEndpoint).Methods("GET")
	// srv.Router.HandleFunc("/weapons/{game}", srv.ReadWeaponsEndpoint).Methods("GET")
	// // returns multiple dmgProfiles
	// srv.Router.HandleFunc("/dmgprofile/{weaponname}", srv.ReadDamageProfiles).Methods("GET")
}
