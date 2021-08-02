package auth

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

type MongoAuth struct {
	URI      string
	Database string
	LoadoutCollection string
	WeaponsCollection string
}

// NewAuth Creates a new MongoAuth struct, containing the username and password necessary to connect to a MongoDB client
// along with the necessary collection names
func NewAuth() *MongoAuth {
	log.Println("Retrieving auth info: " + "func auth.NewAuth()")
	err := godotenv.Load("./config/dev.env")
	if err != nil {
		log.WithFields(log.Fields{
			"func": "NewAuth()",
			"event": "Loading .env file",
			"line": 17,
		}).Fatal(err)
	}

	return &MongoAuth{
		URI:      os.Getenv("MONGO_URI"),
		Database: os.Getenv("DATABASE"),
		LoadoutCollection: os.Getenv("LOADOUTSCOLLECTIONNAME"),
		WeaponsCollection: os.Getenv("WEAPONSCOLLECTIONNAME"),
	}
}
