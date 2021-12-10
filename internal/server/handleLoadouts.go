package server

import (
	"github.com/chancetudor/dubzone-api/internal/models"
	"github.com/chancetudor/dubzone-api/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

// swagger:route POST /loadouts/ loadout newLoadout
// Creates a new loadout.
// responses:
//	200: noContent
// schemes:
//	http, https

/*
CreateLoadout takes a JSON representation of a Loadout,
marshals it into a struct of type Loadout,
and TODO stores that struct in a database.

The function is called only after validation middleware has passed.
TODO change schemes above to just https in prod and add security tag
*/
func (srv *server) CreateLoadout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		srv.log.WithFields(logrus.Fields{"Caller": "CreateLoadout()", "Message": "Creating loadout"}).Info()
		loadout := r.Context().Value(middleware.LoadoutKey{}).(*models.Loadout)
		models.AddProduct(loadout)
	}
}

// swagger:route GET /loadouts/ loadouts listLoadouts
// Returns a list of all loadouts.
// responses:
//	200: loadoutsResponse
// schemes:
//	http, https

// GetLoadouts returns all loadouts in the database in JSON formatting.
func (srv *server) GetLoadouts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		srv.log.WithFields(logrus.Fields{"Caller": "GetLoadouts()", "Message": "Returning all loadouts"}).Info()
		loadouts := models.GetStaticLoadouts()
		w.Header().Add("Content-Type", "application/json")
		err := loadouts.ToJSON(w)
		if err != nil {
			srv.log.Error(errors.Wrap(err, "Unable to marshal JSON into loadouts struct"))
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}

// swagger:route GET /loadouts/meta loadouts listMetaLoadouts
// Returns a list of all loadouts marked as meta.
// responses:
//	200: loadoutsResponse
// schemes:
//	http, https

// GetMetaLoadouts returns all loadouts marked as meta.
func (srv *server) GetMetaLoadouts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		srv.log.WithFields(logrus.Fields{"Caller": "GetMetaLoadouts()", "Message": "Returning all meta loadouts"}).Info()
		meta := models.GetMetaLoadouts()
		w.Header().Add("Content-Type", "application/json")
		err := meta.ToJSON(w)
		if err != nil {
			srv.log.Error(errors.Wrap(err, "Unable to marshal JSON into loadout struct"))
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}

// swagger:route GET /loadouts/category/{category} loadouts listLoadoutsByCategory
// Returns a list of all loadouts whose primary weapon's category matches the category parameter given.
// responses:
//	200: loadoutsResponse
// schemes:
//	http, https

// GetLoadoutsByCategory takes a category parameter and returns all loadouts
// whose primary weapon is tagged with that given category.
// The category parameter is required; if it is not given
// or if the category does not exist, an http.StatusBadRequest is returned.
func (srv *server) GetLoadoutsByCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cat := r.Context().Value(middleware.CatKey{}).(string)
		srv.log.WithFields(logrus.Fields{"Caller": "GetLoadoutsByCategory()", "Message": "Returning all loadouts with category: " + cat}).Info()
		loadouts := models.GetLoadoutsByCategory(cat)
		w.Header().Add("Content-Type", "application/json")
		err := loadouts.ToJSON(w)
		if err != nil {
			srv.log.Error(errors.Wrap(err, "Unable to marshal JSON into loadouts struct"))
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}

// swagger:route GET /loadouts/weapon/{weapon_name} loadouts listLoadoutsByWeapon
// Returns a list of all loadouts whose primary weapon's category matches the name parameter given.
// responses:
//	200: loadoutsResponse
// schemes:
//	http, https

// GetLoadoutsByWeapon takes a name parameter and returns all loadouts
// whose primary weapon is named as such.
// The name parameter is required; if it is not given an http.StatusBadRequest is returned.
func (srv *server) GetLoadoutsByWeapon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.Context().Value(middleware.NameKey{}).(string)
		srv.log.WithFields(logrus.Fields{"Caller": "GetLoadoutsByWeapon()", "Message": "Returning all loadouts with name: " + name}).Info()
		loadouts := models.GetLoadoutsByName(name)
		w.Header().Add("Content-Type", "application/json")
		err := loadouts.ToJSON(w)
		if err != nil {
			srv.log.Error(errors.Wrap(err, "Unable to marshal JSON into loadouts struct"))
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}
