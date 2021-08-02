package api

import (
	"context"
	"github.com/chancetudor/dubzone-api/internal/auth"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type API struct {
	Router *mux.Router
	Client *mongo.Client
}

func NewAPI() *API {
	api := &API{
		Router: newRouter(),
		Client: newClient(),
	}
	api.initRouter()

	return api
}

func newRouter() *mux.Router {
	log.Println("Creating new router: " + "func NewRouter()")
	r := mux.NewRouter().StrictSlash(true)

	return r
}

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
			"func": "NewClient()",
			"event": "Connecting to mongoDB",
		}).Fatal(err)
	}

	return client
}

func (a *API) DisconnectClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			log.WithFields(log.Fields{
				"func": "main()",
				"event": "Client disconnect",
			}).Fatal(err)
		}
	}(a.Client, ctx)
}

// initRouter initializes handler funcs on router
func (a *API) initRouter() {
	log.Println("Initializing router, adding handlers: " + "func InitRouter()")
	// single weapon endpoints, which deal with a single weapon
	a.Router.HandleFunc("/weapon", a.CreateWeaponEndpoint).Methods("POST")
	a.Router.HandleFunc("/weapon/{weaponname}", a.ReadWeaponEndpoint).Methods("GET")
	// single dmgProfile endpoints, which deal with a single dmgProfile for a given weapon
	a.Router.HandleFunc("/dmgprofile/{weaponname}", a.CreateDamageProfileEndpoint).Methods("POST")
	a.Router.HandleFunc("/dmgprofile/{weaponname}", a.ReadDamageProfileEndpoint).Methods("GET")
	a.Router.HandleFunc("/dmgprofile/{weaponname}", a.UpdateDamageProfileEndpoint).Methods("PUT")
	a.Router.HandleFunc("/dmgprofile/{weaponname}", a.DeleteDamageProfileEndpoint).Methods("DELETE")
	// single loadout endpoints, which deal with a single loadout
	a.Router.HandleFunc("/loadout", a.CreateLoadoutEndpoint).Methods("POST")
	// returns multiple weapons
	a.Router.HandleFunc("/weapons", a.ReadWeaponsEndpoint).Methods("GET")
	// returns multiple dmgProfiles
	a.Router.HandleFunc("/dmgprofiles", a.ReadDamageProfilesEndpoint).Methods("GET")
	// returns multiple loadouts
	a.Router.HandleFunc("/loadouts", a.ReadLoadoutsEndpoint).Methods("GET")
}
