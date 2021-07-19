package api

import (
	"context"
	"github.com/chancetudor/dubzone-api/internal/auth"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// The Client file is used to connect directly to the MongoDB database

func NewClient() *mongo.Client {
	log.Println("Creating new client: " + "func NewClient()")
	mongoAuth := auth.New()
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://" + mongoAuth.Username + ":" + mongoAuth.Password +
			"@warzonedata.xbbkl.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	// TODO change timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	// TODO custom error handling
	if err != nil {
		log.Fatal(err)
	}

	return client
}