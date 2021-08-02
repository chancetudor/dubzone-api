package main

import (
	"github.com/chancetudor/dubzone-api/internal/api"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// Version indicates the current version of the application.
var Version = "0.0.1"

func init() {
	// initializes a custom logging file
	file, err := os.OpenFile("./pkg/log/api_logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

func main() {
	log.Debug("Starting application: " + "func main()")
	apiContainer := api.NewAPI()
	defer apiContainer.DisconnectClient()
	log.Debug("Calling ListenAndServe()...")
	err := http.ListenAndServe(":12345", apiContainer.Router)
	if err != nil {
		log.WithFields(log.Fields{
			"func": "main()",
			"event": "ListenAndServe",
		}).Fatal(err)
	}
}
