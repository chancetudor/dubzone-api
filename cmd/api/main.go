package main

import (
	"github.com/chancetudor/dubzone-api/internal/api"
	"log"
)

// Version indicates the current version of the application.
var Version = "0.0.1"

func main() {
	log.Println("Starting application: " + "func main()")
	// create new API struct holding router and client info
	_ = api.New()

}
