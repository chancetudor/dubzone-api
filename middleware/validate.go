package middleware

import (
	"context"
	"github.com/chancetudor/dubzone-api/internal/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Middleware struct {
	Log *logrus.Logger
}

func NewMiddleware(l *logrus.Logger) *Middleware {
	return &Middleware{Log: l}
}

// LoadoutKey is the key to be used for use in middleware.
// When we store the internal representation of a JSON struct,
// we use the context.WithValue function, which takes this as its key.
type LoadoutKey struct{}

/*
ValidateLoadout is a middleware function to determine whether the JSON representation
of type Loadout is valid.

If JSON is valid, ValidateLoadout calls the HandlerFunc it was passed.

If JSON is not valid, ValidateLoadout logs an error and returns http.StatusUnprocessableEntity to the caller.
*/
func (mdl *Middleware) ValidateLoadout(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loadout := &models.Loadout{}
		err := loadout.FromJSON(r.Body)
		if err != nil {
			mdl.Log.Error(errors.Wrap(err, "error deserializing loadout passed as request"))
			http.Error(w, "Error reading product", http.StatusUnprocessableEntity)
			return
		}

		err = loadout.Validate()
		if err != nil {
			mdl.Log.Error(errors.Wrap(err, "error validating loadout passed as request"))
			http.Error(w, "Error validating product: "+err.Error(), http.StatusUnprocessableEntity)
			return
		}
		// add the loadout to the request context for use in next handler
		ctx := context.WithValue(r.Context(), LoadoutKey{}, loadout)
		r = r.WithContext(ctx)
		// call next handler
		h(w, r)
	}
}
