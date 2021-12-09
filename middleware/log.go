package middleware

import (
	"github.com/gorilla/handlers"
	"net/http"
)

// Log utilizes the gorilla handlers package to handle logging of requests for us.
// The function is called as the first middleware function in each chain
// of every request. The outfile is simply the out file of the middleware type's
// logger, the same outfile as the server type's logger.
func (mdl *Middleware) Log(next http.Handler) http.Handler {
	return handlers.LoggingHandler(mdl.log.Out, next)
}
