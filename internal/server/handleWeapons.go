package server

import (
	"github.com/chancetudor/dubzone-api/internal/models"
	"github.com/chancetudor/dubzone-api/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

// swagger:route GET /weapons weapons listWeapons
// Returns a list of all weapons.
// responses:
//	200: weaponsResponse
// schemes:
//	http, https

// GetWeapons returns all weapons in the database in JSON formatting.
func (srv *server) GetWeapons() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		srv.log.WithFields(logrus.Fields{"Caller": "GetWeapons()", "Message": "Returning all weapons"}).Info()
		weapons := models.GetStaticWeapons()
		w.Header().Add("Content-Type", "application/json")
		err := weapons.ToJSON(w)
		if err != nil {
			srv.log.Error(errors.Wrap(err, "Unable to marshal JSON into weapons struct"))
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}

// swagger:route GET /weapons/meta weapons listMetaWeapons
// Returns a list of all weapons marked as meta.
// responses:
//	200: weaponsResponse
// schemes:
//	http, https

// GetMetaWeapons returns all weapons marked as meta.
func (srv *server) GetMetaWeapons() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		srv.log.WithFields(logrus.Fields{"Caller": "GetMetaWeapons()", "Message": "Returning all meta weapons"}).Info()
		meta := models.GetMetaWeapons()
		w.Header().Add("Content-Type", "application/json")
		err := meta.ToJSON(w)
		if err != nil {
			srv.log.Error(errors.Wrap(err, "Unable to marshal JSON into weapon struct"))
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}

// swagger:route GET /weapons/categories weapons listWeaponCategories
// Returns a list of all weapon categories.
// responses:
//	200: categoriesResponse
// schemes:
//	http, https

// GetWeaponCategories returns all weapon categories in the database in JSON formatting.
func (srv *server) GetWeaponCategories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		srv.log.WithFields(logrus.Fields{"Caller": "GetWeaponCategories()", "Message": "Returning all weapon categories"}).Info()
		categories := models.GetWeaponCategories()
		w.Header().Add("Content-Type", "application/json")
		err := categories.ToJSON(w)
		if err != nil {
			srv.log.Error(errors.Wrap(err, "Unable to marshal JSON into categories struct"))
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}

// swagger:route GET /weapons/weapon weapons listWeaponsWithQueryParams
// Returns a list of all weapons that meet a certain parameter (weapon name, category, game) given as a query param.
// responses:
//	200: weaponsResponse
// schemes:
//	http, https

// GetWeaponsByCategory takes a category parameter
// and returns all weapons tagged with that given category.
// The category parameter is required; if it is not given an http.StatusNotFound error is returned.
// If the given category does not exist, an http.StatusBadRequest is returned.
func (srv *server) GetWeaponsByCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cat := r.Context().Value(middleware.CatKey{}).(string)
		srv.log.WithFields(logrus.Fields{"Caller": "GetWeaponsByCategory()", "Message": "Returning all weapons with category: " + cat}).Info()
		weapons := models.GetWeaponsByCategory(cat)
		w.Header().Add("Content-Type", "application/json")
		err := weapons.ToJSON(w)
		if err != nil {
			srv.log.Error(errors.Wrap(err, "Unable to marshal JSON into weapons struct"))
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}

// swagger:route GET /weapons/weapon weapons listWeaponsWithQueryParams
// Returns a list of all weapons that meet a certain parameter (weapon name, category, game) given as a query param.
// responses:
//	200: weaponsResponse
// schemes:
//	http, https

// GetWeaponsByName takes a name parameter and returns all weapon builds
// where the weapon name matches the parameter.
// The name parameter is required; if it is not given an http.StatusNotFound is returned.
func (srv *server) GetWeaponsByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.Context().Value(middleware.NameKey{}).(string)
		srv.log.WithFields(logrus.Fields{
			"Caller":  "GetWeaponsByName()",
			"Message": "Returning all weapons with name: " + name,
		}).Info()
		weapons := models.GetWeaponsByName(name)
		w.Header().Add("Content-Type", "application/json")
		err := weapons.ToJSON(w)
		if err != nil {
			srv.log.Error(errors.Wrap(err, "Unable to marshal JSON into weapons struct"))
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}

// swagger:route GET /weapons/weapon weapons listWeaponsWithQueryParams
// Returns a list of all weapons that meet a certain parameter (weapon name, category, game) given as a query param.
// responses:
//	200: weaponsResponse
// schemes:
//	http, https

// GetWeaponsByGame takes a game parameter and returns
// all weapons where the game matches the parameter.
// The game parameter is required; if it is not given an http.StatusBadRequest is returned.
func (srv *server) GetWeaponsByGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		game := r.Context().Value(middleware.GameKey{}).(string)
		srv.log.WithFields(logrus.Fields{"Caller": "GetWeaponsByGame()",
			"Message": "Returning all weapons with game: " + game}).Info()
		weapons := models.GetWeaponsByGame(game)
		w.Header().Add("Content-Type", "application/json")
		err := weapons.ToJSON(w)
		if err != nil {
			srv.log.Error(errors.Wrap(err, "Unable to marshal JSON into weapons struct"))
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}
