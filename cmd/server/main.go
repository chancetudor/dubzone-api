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
	file, err := os.OpenFile("./log/api_logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Error("Failed to log to file, using default stderr")
	}
	log.ReportCaller = true
}

func main() {
	srv := server.NewServer(log)
	err := srv.Start(":9090")
	if err != nil {
		srv.Log.Fatal(err)
	}
}
