/*
Package models contains the Go representation of our data,
which is passed as JSON to/from the consumer and into our database.

Package models contains Loadout models and Weapon models.
*/

package models

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
	"strings"
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
	Meta *bool `json:"meta_loadout" validate:"required"`
}

// Loadouts describes a slice of type *Loadout
type Loadouts []*Loadout

var validCats map[string]bool = map[string]bool{"RANGE": true, "CLOSE": true, "CLOSE-MED": true, "SNIPER RANGED": true, "SNIPER SUPPORT": true}

// TODO remove as we implement Google Cloud Datastore
var StaticLoadouts = Loadouts{
	&Loadout{Primary: XM4, Secondary: Pistol, PerkOne: "Perk1", PerkTwo: "Perk2", PerkThree: "Perk3", Lethal: "Semtex", Tactical: "Stuns", Meta: &[]bool{true}[0]},
	&Loadout{Primary: C58, Secondary: Pistol, PerkOne: "Perk1", PerkTwo: "Perk2", PerkThree: "Perk3", Lethal: "Semtex", Tactical: "Stuns", Meta: &[]bool{true}[0]},
	&Loadout{Primary: Mac10, Secondary: Pistol, PerkOne: "Perk1", PerkTwo: "Perk2", PerkThree: "Perk3", Lethal: "Semtex", Tactical: "Stuns", Meta: &[]bool{true}[0]},
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
		if l.Meta == &[]bool{true}[0] {
			MetaLoadouts = append(MetaLoadouts, l)
		}
	}

	return MetaLoadouts
}

// TODO remove as we get Google Cloud Datastore in operation
func GetLoadoutsByCategory(cat string) Loadouts {
	CatLoadouts := Loadouts{}
	for _, l := range StaticLoadouts {
		if strings.EqualFold(cat, l.Primary.Category) {
			CatLoadouts = append(CatLoadouts, l)
		}
	}

	return CatLoadouts
}

func GetLoadoutsByName(name string) Loadouts {
	loadouts := Loadouts{}
	for _, l := range StaticLoadouts {
		if strings.EqualFold(name, l.Primary.WeaponName) {
			loadouts = append(loadouts, l)
		}
	}

	return loadouts
}

// ValidCategory is called in ValidateCategoryParam and returns a bool representing whether
// the category parameter is valid or not. If the category passed is "snipersupport" or "sniperranged,"
// the function transforms that string into "sniper support" and/or "sniper ranged," as these
// are the correct values. The caller is not expected to know this.
func ValidCategory(cat string) (string, bool) {
	validCat := transformCategory(cat)
	_, valid := validCats[strings.ToUpper(validCat)]
	if valid {
		return validCat, true
	}
	return "", false
}

func transformCategory(c string) string {
	switch {
	case strings.EqualFold(c, "snipersupport"):
		return "sniper support"
	case strings.EqualFold(c, "sniperranged"):
		return "sniper ranged"
	}

	return c
}