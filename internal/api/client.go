package api

import (
	"context"
	"github.com/chancetudor/dubzone-api/internal/auth"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// The Client file is used to connect directly to the MongoDB database

func NewClient() *mongo.Client {
	log.Println("Creating new client: " + "func NewClient()")
	mongoAuth := auth.NewAuth()
	clientOptions := options.Client().
		ApplyURI(mongoAuth.URI)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	// TODO custom error handling
	if err != nil {
		log.WithFields(log.Fields{
			"func": "NewClient()",
			"event": "Connecting to mongoDB",
			"line": 21,
		}).Fatal(err)
	}

	return client
}
