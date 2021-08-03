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
	Auth *auth.MongoAuth
}

// NewAPI is a constructor for our API struct
func NewAPI() *API {
	api := &API{
		Router: newRouter(),
		Client: newClient(),
		Auth: auth.NewAuth(),
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
			"func": "NewClient()",
			"event": "Connecting to mongoDB",
		}).Fatal(err)
	}

	return client
}

// DisconnectClient disconnects the mongo client
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
	// creates a single new weapon in Weapons collection
	a.Router.HandleFunc("/weapon", a.CreateWeaponEndpoint).Methods("POST")
	// gets a specified weapon from Weapons collection
	a.Router.HandleFunc("/weapon/{weaponname}", a.ReadWeaponEndpoint).Methods("GET")
	// gets all damage profiles for specified weapon
	a.Router.HandleFunc("/dmgprofile/{weaponname}", a.ReadDamageProfileEndpoint).Methods("GET")
	// gets close-range damage profiles for specified weapon
	a.Router.HandleFunc("/dmgprofile/{weaponname}/close", a.ReadCloseDamageProfileEndpoint).Methods("GET")
	// gets mid-range damage profiles for specified weapon
	a.Router.HandleFunc("/dmgprofile/{weaponname}/mid", a.ReadMidDamageProfileEndpoint).Methods("GET")
	// gets far-range damage profiles for specified weapon
	a.Router.HandleFunc("/dmgprofile/{weaponname}/far", a.ReadFarDamageProfileEndpoint).Methods("GET")
	// updates a close-range damage profile for a specified weapon
	a.Router.HandleFunc("/dmgprofile/{weaponname}/close", a.UpdateCloseDamageProfileEndpoint).Methods("PUT")
	// updates a mid-range damage profile for a specified weapon
	a.Router.HandleFunc("/dmgprofile/{weaponname}/mid", a.UpdateMidDamageProfileEndpoint).Methods("PUT")
	// updates a close-range damage profile for a specified weapon
	a.Router.HandleFunc("/dmgprofile/{weaponname}/far", a.UpdateFarDamageProfileEndpoint).Methods("PUT")
	// creates a single new loadout in Loadouts collection
	a.Router.HandleFunc("/loadout", a.CreateLoadoutEndpoint).Methods("POST")
	// returns multiple weapons
	a.Router.HandleFunc("/weapons", a.ReadWeaponsEndpoint).Methods("GET")
	a.Router.HandleFunc("/weapons/{game}", a.ReadWeaponsByGameEndpoint).Methods("GET")
	// returns multiple loadouts
	a.Router.HandleFunc("/loadouts", a.ReadLoadoutsEndpoint).Methods("GET")
	a.Router.HandleFunc("/loadouts/{category}", a.ReadLoadoutsByCategoryEndpoint).Methods("GET")
	a.Router.HandleFunc("/loadouts/{weaponname}", a.ReadLoadoutsByWeaponEndpoint).Methods("GET")
}
