package main

import (
	"github.com/chancetudor/dubzone-api/internal/server"
	"github.com/sirupsen/logrus"
	"os"
)

// Version indicates the current version of the application.
var Version = "0.0.2"
var log = logrus.New()

func init() {
	// TODO set to prod_logs when in production
	file, err := os.OpenFile("./logs/test_logs.logs", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error("Failed to logs to file, using default stderr")
		return

	}
	log.SetOutput(file)
	log.Formatter = &logrus.JSONFormatter{PrettyPrint: true}
}

func main() {
	srv := server.NewServer(log)
	srv.Start(":9090")
}
