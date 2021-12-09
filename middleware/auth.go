package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func (mdl *Middleware) Authenticate(next http.Handler) http.Handler {
	mdl.log.WithFields(logrus.Fields{"Caller": "Authenticate()", "Message": "Not implemented"}).Info()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
