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
	Body models.Loadouts
}

// swagger:response weaponsResponse
type weaponsResponse struct {
	// All loadouts in the database.
	// in: body
	Body models.Weapons
}

// swagger:response categoriesResponse
type categoriesResponse struct {
	// All categories in the database.
	// in: body
	Body models.Categories
}

// swagger:parameters listLoadoutsWithQueryParams
type listLoadoutsQueryParams struct {
	//
	// The category by which to return loadouts.
	// in: query
	// pattern: [a-zA-Z]+[-]?[a-zA-Z]*
	// required: true
	Category string `json:"category"`
	// The name of the weapon by which to return loadouts.
	// in: query
	// pattern: [a-zA-Z]+\d*
	// required: true
	WeaponName string `json:"name"`
	// The name of the game by which to return loadout configurations.
	// in: query
	// pattern: [a-zA-Z]+[\s]?[a-zA-Z]+
	// required: true
	Game string `json:"game"`
}

// swagger:parameters listWeaponsWithQueryParams
type listWeaponsQueryParams struct {
	// These are the query parameters to pass when using a /weapons/weapon endpoint.
	// All are marked as required, but this means that you must use one and only one query parameter with the endpoint.
	//
	// The category by which to return weapon configurations.
	// in: query
	// pattern: [a-zA-Z]+[-]?[a-zA-Z]*
	// required: true
	Category string `json:"category"`
	// The name of the weapon by which to return weapon configurations.
	// in: query
	// pattern: [a-zA-Z]+\d*
	// required: true
	WeaponName string `json:"name"`
	// The name of the game by which to return weapon configurations.
	// in: query
	// pattern: [a-zA-Z]+[\s]?[a-zA-Z]+
	// required: true
	Game string `json:"game"`
}
