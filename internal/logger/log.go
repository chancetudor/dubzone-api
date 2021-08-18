package logger

import (
	errs "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
)

// InitLogger performs necessary initialization; at this point, tying a log file to logrus
func InitLogger() {
	// initializes a custom logging file
	file, err := os.OpenFile("./pkg/log/api_logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Error("Failed to log to file, using default stderr")
	}

	log.Debug("STARTING APPLICATION")
}

// Fatal is a helper function to abstract fatal err handling and logging
func Fatal(err error, msg string, function string) {
	log.WithFields(log.Fields{
		"func":  function,
		"event": msg,
	}).Fatal(err)
}

// Error is a helper function to abstract err handling and logging
func Error(err error, msg string, function string) {
	errs.Errorf(msg, err)
	log.WithFields(log.Fields{
		"func":  function,
		"event": msg,
	}).Error(err)
}

// Debug is a helper function to abstract debug information
func Debug(msg string, function string) {
	log.WithFields(log.Fields{
		"func": function,
		"event": msg,
	}).Debug()
}

func clearLogs() {
	if err := os.Truncate("./pkg/log/api_logs.log", 0); err != nil {
		log.Error("Failed to truncate: %v", err)
	}
}