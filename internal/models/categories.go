package models

import (
	"encoding/json"
	"io"
)

type Category struct {
	Category string
}

type Categories []Category

// TODO ensure we have the right categories
var StaticCategories = Categories{
	Category{Category: "Range"},
	Category{Category: "Close"},
	Category{Category: "Close-Med"},
	Category{Category: "Sniper Support"},
	Category{Category: "Sniper Ranged"},
	Category{Category: "Secondary"},
	Category{Category: "Non-overkill"},
}

// FromJSON takes in an io.Reader, the *http.Request body,
// and unmarshals that body into a Loadout.
func (l *Categories) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(l)
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (l *Categories) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(l)
}

func GetWeaponCategories() Categories {
	return StaticCategories
}
