package server

import (
	"context"
	"github.com/chancetudor/dubzone-api/internal/auth"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	r := mux.NewRouter().StrictSlash(true)

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
	// a.Router.HandleFunc("/dmgprofile/{weaponname}", a.CreateDamageProfileEndpoint).Methods("POST")
	srv.Router.HandleFunc("/dmgprofile/{weaponname}", srv.ReadDamageProfileEndpoint).Methods("GET")
	srv.Router.HandleFunc("/dmgprofile/{weaponname}", srv.UpdateDamageProfileEndpoint).Methods("PUT")
	// a.Router.HandleFunc("/dmgprofile/{weaponname}", a.DeleteDamageProfileEndpoint).Methods("DELETE")
	// single loadout endpoints, which deal with a single loadout
	srv.Router.HandleFunc("/loadout", srv.CreateLoadoutEndpoint).Methods("POST")
	// returns multiple weapons
	srv.Router.HandleFunc("/weapons", srv.ReadWeaponsEndpoint).Methods("GET")
	// returns multiple dmgProfiles
	srv.Router.HandleFunc("/dmgprofiles", srv.ReadDamageProfilesEndpoint).Methods("GET")
	// returns multiple loadouts
	srv.Router.HandleFunc("/loadouts", srv.ReadLoadoutsEndpoint).Methods("GET")
}
