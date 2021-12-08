/*
This file contains all swagger responses.
These types are not used for any of our handlers.
*/

package openapi

import "github.com/chancetudor/dubzone-api/internal/models"

// swagger:response noContent
type noContent struct{}

// swagger:response loadoutsResponse
type loadoutsResponse struct {
	// All loadouts in the database.
	// in: body
	Body []*models.Loadout
}

// swagger:parameters listLoadoutsByCategory
type loadoutsByCategoryParameter struct {
	// The category by which to return loadouts.
	// in: path
	// required: true
	Category string `json:"category"`
}

// swagger:parameters listWeaponsByCategory
type weaponsByCategoryParameter struct {
	// The category by which to return weapons.
	// in: path
	// required: true
	Category string `json:"category"`
}

// swagger:parameters listLoadoutsByWeapon
type loadoutsByWeaponParameter struct {
	// The name of the weapon by which to return loadouts.
	// in: path
	// required: true
	Name string `json:"weapon_name"`
}

// swagger:parameters listWeaponsByName
type weaponsByNameParameter struct {
	// The name of the weapon by which to return weapon configurations.
	// in: path
	// required: true
	Name string `json:"weapon_name"`
}
