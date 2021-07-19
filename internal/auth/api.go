package auth

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type MongoAuth struct {
	Username string
	Password string
}

func New() *MongoAuth {
	log.Println("Retrieving auth info: " + "func auth.New()")
	err := godotenv.Load("./config/dev.env")
	// TODO custom error handling
	if err != nil {
		log.Fatal(err)
	}

	return &MongoAuth {
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
}
