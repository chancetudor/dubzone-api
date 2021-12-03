package server

import (
	"os"
)

func (srv *server) clearLogs() {
	if err := os.Truncate("./pkg/Log/api_logs.Log", 0); err != nil {
		srv.Log.Error("Failed to truncate: %v", err)
	}
}
