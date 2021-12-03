package server

import "github.com/chancetudor/dubzone-api/internal/models"

// A list of loadouts returned in a successful response.
// swagger:response loadoutsResponse
type loadoutsResponse struct {
	// All loadouts in the database.
	// in: body
	Body []*models.Loadout
}
