package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

// TODO implement semi-custom authentication
func (mdl *Middleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mdl.log.WithFields(logrus.Fields{"Caller": "Authenticate()", "Message": "Not implemented"}).Info()
		next.ServeHTTP(w, r)
	})
}
