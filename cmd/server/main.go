package main

import (
	"github.com/chancetudor/dubzone-api/internal/logger"
	"github.com/chancetudor/dubzone-api/internal/server"
	"net/http"
)

// Version indicates the current version of the application.
var Version = "0.0.1"

func main() {
	logger.InitLogger()
	srv := server.NewServer()
	defer srv.DisconnectClient()
	logger.Debug("Calling ListenAndServe()...", "main()")
	err := http.ListenAndServe(":12345", srv.Router)
	if err != nil {
		logger.Error(err, "ListenAndServe", "main()")
	}
}
