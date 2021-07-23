package main

import (
	"github.com/chancetudor/dubzone-api/internal/api"
	"log"
	"net/http"
)

// Version indicates the current version of the application.
var Version = "0.0.1"

func main() {
	log.Println("Starting application: " + "func main()")
	dubzoneRouter := api.NewRouter()
	api.InitRouter(dubzoneRouter)
	log.Println("Calling ListenAndServe()...")
	err := http.ListenAndServe(":12345", dubzoneRouter)
	if err != nil {
		log.Fatal(err)
	}
}
