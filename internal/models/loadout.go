package models

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
)

// Loadout represents a Warzone loadout, complete with a primary weapon, secondary weapon, three perks, and lethal and tactical equipment
type Loadout struct {
	// ID        p.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Primary   Weapon `json:"primary,omitempty" validate:"required, dive"`
	Secondary Weapon `json:"secondary,omitempty" validate:"required, dive"`
	PerkOne   string `json:"perk_one,omitempty" validate:"required, alpha, max=15"`
	PerkTwo   string `json:"perk_two,omitempty" validate:"required, alpha, max=15"`
	PerkThree string `json:"perk_three,omitempty" validate:"required, alpha, max=15"`
	Lethal    string `json:"lethal,omitempty" validate:"required, alpha, max=15"`
	Tactical  string `json:"tactical,omitempty" validate:"required, alpha, max=15"`
	Meta      bool   `json:"meta_loadout,omitempty" validate:"required, boolean"`
}

// Loadouts describes a slice of type *Loadout
type Loadouts []*Loadout

// TODO remove as we implement Google Cloud Datastore
var StaticLoadouts = Loadouts{
	&Loadout{Primary: XM4, Secondary: Pistol, PerkOne: "Perk1", PerkTwo: "Perk2", PerkThree: "Perk3", Lethal: "Semtex", Tactical: "Stuns"},
	&Loadout{Primary: C58, Secondary: Pistol, PerkOne: "Perk1", PerkTwo: "Perk2", PerkThree: "Perk3", Lethal: "Semtex", Tactical: "Stuns"},
	&Loadout{Primary: Mac10, Secondary: Pistol, PerkOne: "Perk1", PerkTwo: "Perk2", PerkThree: "Perk3", Lethal: "Semtex", Tactical: "Stuns"},
}

// FromJSON takes in an io.Reader, the *http.Request body,
// and unmarshals that body into a Loadout.
func (l *Loadout) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(l)
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (l *Loadouts) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(l)
}

func (l *Loadout) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

// TODO remove as we get Google Cloud Datastore in operation
func GetStaticLoadouts() Loadouts {
	return StaticLoadouts
}

// TODO remove as we get Google Cloud Datastore in operation
func AddProduct(l *Loadout) {
	StaticLoadouts = append(StaticLoadouts, l)
}
