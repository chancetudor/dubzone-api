package middleware

import (
	"context"
	"fmt"
	"github.com/chancetudor/dubzone-api/internal/models"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"mime"
	"net/http"
)

// func (mdl *Middleware) ValidateLoadout(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		contentType := r.Header.Get("Content-Type")
//
// 		if contentType == "" {
// 			err := errors.New("Content-Type header is empty")
// 			mdl.logs.Error(fmt.Sprintf("%+v", err))
// 			http.Error(w, "Content-Type header is empty", http.StatusBadRequest)
// 			return
// 		}
// 		mt, _, err := mime.ParseMediaType(contentType)
// 		if err != nil {
// 			http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
// 			return
// 		}
//
// 		if mt != "application/json" {
// 			http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
// 			return
// 		}
//
// 		loadout := &models.Loadout{}
// 		err = loadout.FromJSON(r.Body)
// 		if err != nil {
// 			wrappedErr := errors.Wrap(err, "error deserializing loadout passed as request")
// 			mdl.logs.Error(fmt.Sprintf("%+v", wrappedErr))
// 			http.Error(w, "Error reading product", http.StatusUnprocessableEntity)
// 			return
// 		}
//
// 		err = loadout.Validate()
// 		if err != nil {
// 			wrappedErr := errors.Wrap(err, "error validating loadout passed as request")
// 			mdl.logs.Error(fmt.Sprintf("%+v", wrappedErr))
// 			http.Error(w, "Error validating product: "+err.Error(), http.StatusUnprocessableEntity)
// 			return
// 		}
// 		// add the loadout to the request context for use in next handler
// 		ctx := context.WithValue(r.Context(), LoadoutKey{}, loadout)
// 		r = r.WithContext(ctx)
// 		// call next handler
// 		h(w, r)
// 	}
// }

// func (mdl *Middleware) ValidateCategoryParam(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		cat, found := mux.Vars(r)["category"]
// 		if !found {
// 			err := errors.New("no category passed as a parameter")
// 			mdl.logs.Error(fmt.Sprintf("%+v", err))
// 			http.Error(w, "Error validating parameter: no category parameter was passed", http.StatusBadRequest)
// 			return
// 		}
// 		validCat, valid := models.ValidCategory(cat)
// 		if !valid {
// 			err := errors.New("category parameter: " + cat + " was not found")
// 			mdl.logs.Error(fmt.Sprintf("%+v", err))
// 			http.Error(w, "Error validating parameter: incorrect category parameter was passed. "+
// 				"Use /weapons/categories endpoint to find acceptable categories.", http.StatusBadRequest)
// 			return
// 		}
// 		// add the category to the request context for use in next handler
// 		ctx := context.WithValue(r.Context(), CatKey{}, validCat)
// 		r = r.WithContext(ctx)
// 		h(w, r)
// 	}
// }

// func (mdl *Middleware) ValidateNameParam(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		name, found := mux.Vars(r)["name"]
// 		if !found {
// 			err := errors.New("no category passed as a parameter")
// 			mdl.logs.Error(fmt.Sprintf("%+v", err))
// 			http.Error(w, "Error validating parameter: no category parameter was passed", http.StatusBadRequest)
// 			return
// 		}
// 		// add the category to the request context for use in next handler
// 		ctx := context.WithValue(r.Context(), NameKey{}, name)
// 		r = r.WithContext(ctx)
// 		h(w, r)
// 	}
// }

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
			mdl.log.Error(fmt.Sprintf("%+v", err))
			http.Error(w, "Content-Type header is empty", http.StatusBadRequest)
			return
		}
		mt, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
			return
		}

		if mt != "application/json" {
			http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
			return
		}

		loadout := &models.Loadout{}
		err = loadout.FromJSON(r.Body)
		if err != nil {
			wrappedErr := errors.Wrap(err, "error deserializing loadout passed as request")
			mdl.log.Error(fmt.Sprintf("%+v", wrappedErr))
			http.Error(w, "Error reading product", http.StatusUnprocessableEntity)
			return
		}

		err = loadout.Validate()
		if err != nil {
			wrappedErr := errors.Wrap(err, "error validating loadout passed as request")
			mdl.log.Error(fmt.Sprintf("%+v", wrappedErr))
			http.Error(w, "Error validating product: "+err.Error(), http.StatusUnprocessableEntity)
			return
		}
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
*/
func (mdl *Middleware) ValidateCategoryParam(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mdl.log.WithFields(logrus.Fields{"Caller": "ValidateCategoryParam()", "Message": "Validating category parameter"}).Info()
		cat, found := mux.Vars(r)["category"]
		if !found {
			err := errors.New("no category passed as a parameter")
			mdl.log.Error(fmt.Sprintf("%+v", err))
			http.Error(w, "Error validating parameter: no category parameter was passed", http.StatusBadRequest)
			return
		}
		validCat, valid := models.ValidCategory(cat)
		if !valid {
			err := errors.New("category parameter: " + cat + " was not found")
			mdl.log.Error(fmt.Sprintf("%+v", err))
			http.Error(w, "Error validating parameter: incorrect category parameter was passed. "+
				"Use /weapons/categories endpoint to find acceptable categories.", http.StatusBadRequest)
			return
		}
		// add the category to the request context for use in next handler
		ctx := context.WithValue(r.Context(), CatKey{}, validCat)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

/*
ValidateNameParam is a middleware function to determine whether the weapon name parameter is valid.

If parameter is valid, ValidateNameParam calls the HandlerFunc it was passed, adding the category (after possible
transformation) to the request context.

If parameter is not valid, ValidateNameParam logs an error and returns http.StatusBadRequest to the caller.
*/
func (mdl *Middleware) ValidateNameParam(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mdl.log.WithFields(logrus.Fields{"Caller": "ValidateNameParam()", "Message": "Validating name parameter"}).Info()
		name, found := mux.Vars(r)["weapon_name"]
		mdl.log.WithFields(logrus.Fields{"Caller": "ValidateNameParam()", "NameParam": name}).Info()
		if !found {
			err := errors.New("no name passed as a parameter")
			mdl.log.Error(fmt.Sprintf("%+v", errors.Unwrap(err)))
			http.Error(w, "Error validating parameter: no category parameter was passed", http.StatusBadRequest)
			return
		}
		// add the category to the request context for use in next handler
		ctx := context.WithValue(r.Context(), NameKey{}, name)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
