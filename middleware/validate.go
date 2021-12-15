package middleware

import (
	"context"
	"github.com/chancetudor/dubzone-api/internal/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"mime"
	"net/http"
)

/*
ValidateLoadout is a middleware function to determine whether the JSON representation
of type Loadout is valid.

If JSON is valid, ValidateLoadout calls the HandlerFunc it was passed.

If JSON is not valid, ValidateLoadout logs an error and returns http.StatusUnprocessableEntity to the caller.
*/
func (mdl *Middleware) ValidateLoadout(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType == "" {
			err := errors.New("Content-Type header is empty")
			mdl.log.Error(err)
			http.Error(w, "Content-Type header is empty", http.StatusBadRequest)
			return
		}
		parsedContentType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
			return
		}

		if parsedContentType != "application/json" {
			http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
			return
		}

		loadout := &models.Loadout{}
		err = loadout.FromJSON(r.Body)
		if err != nil {
			wrappedErr := errors.Wrap(err, "error deserializing loadout passed as request")
			mdl.log.Error(wrappedErr)
			http.Error(w, "Error reading product", http.StatusUnprocessableEntity)
			return
		}

		err = loadout.Validate()
		if err != nil {
			wrappedErr := errors.Wrap(err, "error validating loadout passed as request")
			mdl.log.Error(wrappedErr)
			http.Error(w, "Error validating product: "+err.Error(), http.StatusUnprocessableEntity)
			return
		}
		// TODO loadout.Clean() --> ensure each field is ALL CAPS and strippedofspaces
		// add the loadout to the request context for use in next handler
		ctx := context.WithValue(r.Context(), LoadoutKey{}, loadout)
		r = r.WithContext(ctx)
		// call next handler
		next.ServeHTTP(w, r)
	})
}

/*
ValidateCategoryParam is a middleware function to determine whether the category parameter is valid.

If parameter is valid, ValidateCategoryParam calls the HandlerFunc it was passed, adding the category (after possible
transformation) to the request context.

If parameter is not valid, ValidateCategoryParam logs an error and returns http.StatusBadRequest to the caller.
If parameter does not exist, http.StatusBadRequest is returned.
*/
func (mdl *Middleware) ValidateCategoryParam(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mdl.log.WithFields(logrus.Fields{"Caller": "ValidateCategoryParam()", "Message": "Validating category parameter"}).Info()
		cat := r.FormValue("category")
		if cat == "" {
			err := errors.New("no category passed as a parameter")
			mdl.log.Error(err)
			http.Error(w, "Error validating parameter: no category parameter was passed", http.StatusBadRequest)
			return
		}
		formattedCat, valid := models.ValidCategory(cat)
		if !valid {
			err := errors.New("category parameter: " + cat + " was not found")
			mdl.log.Error(err)
			http.Error(w, "Error validating parameter: incorrect category parameter was passed. "+
				"Use /weapons/categories endpoint to find acceptable categories.", http.StatusBadRequest)
			return
		}
		// add the category to the request context for use in next handler
		ctx := context.WithValue(r.Context(), CatKey{}, formattedCat)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

/*
ValidateWeaponNameParam is a middleware function to determine whether the weapon name parameter is valid.

If parameter is valid, ValidateWeaponNameParam calls the HandlerFunc it was passed, adding the category (after possible
transformation) to the request context.

If parameter is not valid, ValidateWeaponNameParam logs an error and returns http.StatusBadRequest to the caller.
If parameter does not exist, http.StatusBadRequest is returned.
*/
func (mdl *Middleware) ValidateWeaponNameParam(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		if name == "" {
			err := errors.New("no name passed as a parameter")
			mdl.log.Error(err)
			http.Error(w, "Error validating parameter: no name parameter was passed", http.StatusBadRequest)
			return
		}
		mdl.log.WithFields(logrus.Fields{"Caller": "ValidateWeaponNameParam()", "NameParam": name}).Info()
		// add the category to the request context for use in next handler
		ctx := context.WithValue(r.Context(), NameKey{}, name)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

/*
ValidateGameParam is a middleware function to determine whether the weapon name parameter is valid.

If parameter is valid, ValidateGameParam calls the HandlerFunc it was passed, adding the category (after possible
transformation) to the request context.

If parameter is not valid, ValidateGameParam logs an error and returns http.StatusBadRequest to the caller.
If parameter does not exist, http.StatusBadRequest is returned.
*/
func (mdl *Middleware) ValidateGameParam(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mdl.log.WithFields(logrus.Fields{"Caller": "ValidateGameParam()", "Message": "Validating game parameter"}).Info()
		game := r.FormValue("game")
		mdl.log.WithFields(logrus.Fields{"Caller": "ValidateGameParam()", "GameParam": game}).Info()
		if game == "" {
			err := errors.New("no game passed as a parameter")
			mdl.log.Error(err)
			http.Error(w, "Error validating parameter: no game parameter was passed", http.StatusBadRequest)
			return
		}
		formattedGame, valid := models.ValidGame(game)
		if !valid {
			err := errors.New("game parameter: " + game + " was not found")
			mdl.log.Error(err)
			http.Error(w, "Error validating parameter: incorrect game parameter was passed.", http.StatusBadRequest)
			return
		}
		// add the category to the request context for use in next handler
		ctx := context.WithValue(r.Context(), GameKey{}, formattedGame)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
