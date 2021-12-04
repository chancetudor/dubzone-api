package server

import (
	"github.com/chancetudor/dubzone-api/internal/models"
	"github.com/chancetudor/dubzone-api/middleware"
	"github.com/pkg/errors"
	"net/http"
)

// swagger:route POST /loadouts loadout newLoadout
// Creates a new loadout.
// responses:
//	200: noContent

/*
CreateLoadout takes a JSON representation of a Loadout,
marshals it into a struct of type Loadout,
and TODO stores that struct in a database.

The function is called only after validation middleware has passed.
*/
func (srv *server) CreateLoadout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loadout := r.Context().Value(middleware.LoadoutKey{}).(*models.Loadout)
		models.AddProduct(loadout)
	}
}

// swagger:route GET /loadouts loadouts listLoadouts
// Returns a list of all loadouts.
// responses:
//	200: loadoutsResponse

// GetLoadouts returns all loadouts in the database in JSON formatting.
func (srv *server) GetLoadouts() http.HandlerFunc {
	srv.Log.Info("Getting loadouts")
	return func(w http.ResponseWriter, r *http.Request) {
		loadouts := models.GetStaticLoadouts()
		w.Header().Add("Content-Type", "application/json")
		err := loadouts.ToJSON(w)
		if err != nil {
			srv.Log.Error(errors.Wrap(err, "Unable to marshal JSON into loadout struct"))
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}

// swagger:route GET /loadouts/meta loadouts listMetaLoadouts
// Returns a list of all loadouts marked as meta.
// responses:
//	200: loadoutsResponse

// GetMetaLoadouts returns all loadouts marked as meta.
func (srv *server) GetMetaLoadouts() http.HandlerFunc {
	srv.Log.Info("Getting meta loadouts")
	return func(w http.ResponseWriter, r *http.Request) {
		meta := models.GetMetaLoadouts()
		w.Header().Add("Content-Type", "application/json")
		err := meta.ToJSON(w)
		if err != nil {
			srv.Log.Error(errors.Wrap(err, "Unable to marshal JSON into loadout struct"))
			http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		}
	}
}
