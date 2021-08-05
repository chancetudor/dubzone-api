package server

import (
	"context"
	"encoding/json"
	"github.com/chancetudor/dubzone-api/internal/auth"
	"github.com/gorilla/mux"
	errs "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
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
	log.Println("Creating new router: " + "func NewRouter()")
	r := mux.NewRouter().StrictSlash(true) //.UseEncodedPath() TODO add in and unescape paramters where necesse est

	return r
}

// newClient creates a new mongo client with appropriate authentication
func newClient() *mongo.Client {
	log.Println("Creating new client: " + "func NewClient()")

	mongoAuth := auth.NewAuth()
	clientOptions := options.Client().
		ApplyURI(mongoAuth.URI)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.WithFields(log.Fields{
			"func":  "NewClient()",
			"event": "Connecting to mongoDB",
		}).Fatal(err)
	}

	return client
}

// error is a helper function to abstract err handling and logging
func (srv *server) error(err error, msg string) {
	errs.Errorf(msg, err)
	log.Error(err.Error(), msg)
}

// respond is a helper function to abstract HTTP responses
func (srv *server) respond(response http.ResponseWriter, data interface{}, status int) {
	response.Header().Add("content-type", "application/json")
	response.WriteHeader(status)

	if data == nil {
		response.Write([]byte(`{"message": "Error retrieving data"}`))
		return
	}

	if data != nil {
		if err := json.NewEncoder(response).Encode(data); err != nil {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(`{"message": "` + err.Error() + `"}`))
			log.WithFields(log.Fields{
				"func":  "srv.response()",
				"event": "Encoding data into JSON response",
			}).Error(err)
		}
	}
}

func (srv *server) DisconnectClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			log.WithFields(log.Fields{
				"func":  "main()",
				"event": "Client disconnect",
			}).Fatal(err)
		}
	}(srv.Client, ctx)
}

// initRouter initializes handler funcs on router
func (srv *server) initRouter() {
	log.Println("Initializing router, adding handlers: " + "func InitRouter()")
	// single weapon endpoints, which deal with a single weapon
	srv.Router.HandleFunc("/weapon", srv.CreateWeaponEndpoint).Methods("POST")
	srv.Router.HandleFunc("/weapon/{weaponname}", srv.ReadWeaponEndpoint).Methods("GET")
	// single dmgProfile endpoints, which deal with a single dmgProfile for a given weapon
	srv.Router.HandleFunc("/dmgprofile/{weaponname}", srv.ReadDamageProfileEndpoint).Methods("GET")
	srv.Router.HandleFunc("/dmgprofile/{weaponname}", srv.UpdateDamageProfileEndpoint).Methods("PUT")
	// single loadout endpoints, which deal with a single loadout
	srv.Router.HandleFunc("/loadout", srv.CreateLoadoutEndpoint).Methods("POST")
	// returns multiple weapons
	srv.Router.HandleFunc("/weapons", srv.ReadWeaponsEndpoint).Methods("GET")
	srv.Router.HandleFunc("/weapons/{game}", srv.ReadWeaponsEndpoint).Methods("GET")
	// returns multiple dmgProfiles
	srv.Router.HandleFunc("/dmgprofiles", srv.ReadDamageProfilesEndpoint).Methods("GET")
	// returns multiple loadouts
	srv.Router.HandleFunc("/loadouts", srv.ReadLoadoutsEndpoint).Methods("GET")
	srv.Router.HandleFunc("/loadouts/{category}", srv.ReadLoadoutsEndpoint).Methods("GET")
	srv.Router.HandleFunc("/loadouts/{weaponname}", srv.ReadLoadoutsEndpoint).Methods("GET")
}
