package models

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
)

// Loadout represents a Warzone loadout
//
// A Loadout is the principal object for this service
//
// Loadouts are comprised of two Weapons, three perks, and tactical and lethal equipment
// swagger:model
type Loadout struct {
	// ID        p.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	// The primary weapon for this loadout
	//
	// required: true
	Primary Weapon `json:"primary" validate:"required"`
	// The secondary weapon for this loadout
	//
	// required: true
	Secondary Weapon `json:"secondary" validate:"required"`
	// The first perk slot for this loadout
	//
	// required: false
	PerkOne string `json:"perk_one,omitempty" validate:"omitempty,ascii,max=25"`
	// The second perk slot for this loadout
	//
	// required: false
	PerkTwo string `json:"perk_two,omitempty" validate:"omitempty,ascii,max=25"`
	// The third perk slot for this loadout
	//
	// required: false
	PerkThree string `json:"perk_three,omitempty" validate:"omitempty,ascii,max=25"`
	// The lethal equipment for this loadout
	//
	// required: false
	Lethal string `json:"lethal,omitempty" validate:"omitempty,ascii,max=15"`
	// The tactical equipment for this loadout
	//
	// required: false
	Tactical string `json:"tactical,omitempty" validate:"omitempty,ascii,max=18"`
	// Marks whether this is a meta loadout or not
	//
	// required: true
	// example: true
	Meta bool `json:"meta_loadout" validate:"required"`
}

// Loadouts describes a slice of type *Loadout
type Loadouts []*Loadout

// TODO remove as we implement Google Cloud Datastore
var StaticLoadouts = Loadouts{
	&Loadout{Primary: XM4, Secondary: Pistol, PerkOne: "Perk1", PerkTwo: "Perk2", PerkThree: "Perk3", Lethal: "Semtex", Tactical: "Stuns", Meta: true},
	&Loadout{Primary: C58, Secondary: Pistol, PerkOne: "Perk1", PerkTwo: "Perk2", PerkThree: "Perk3", Lethal: "Semtex", Tactical: "Stuns", Meta: true},
	&Loadout{Primary: Mac10, Secondary: Pistol, PerkOne: "Perk1", PerkTwo: "Perk2", PerkThree: "Perk3", Lethal: "Semtex", Tactical: "Stuns", Meta: false},
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

// TODO remove as we get Google Cloud Datastore in operation
func GetMetaLoadouts() Loadouts {
	MetaLoadouts := Loadouts{}
	for _, l := range StaticLoadouts {
		if l.Meta == true {
			MetaLoadouts = append(MetaLoadouts, l)
		}
	}

	return MetaLoadouts
}
