package models

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
	"strings"
)

// Weapon represents a Warzone weapon, complete with a category and all recommended attachments.
// A Weapon can have a maximum of 5 fields set at one time. TODO potentially change w/ Vanguard integration
// Leave the field as an empty string ("") if the field is not to be set.
type Weapon struct {
	// The weapon's name
	//
	// required: true
	WeaponName string `json:"weapon_name" validate:"required,ascii"`
	// The game the weapon is from
	//
	// required: true
	Game string `json:"game" validate:"required,ascii"`
	// The weapon's category
	//
	// required: true
	// example: Range
	Category string `json:"category" validate:"required,ascii"`
	// The weapon's muzzle attachment
	//
	// required: false
	Muzzle string `json:"muzzle,omitempty" validate:"omitempty,ascii"`
	// The weapon's barrel attachment
	//
	// required: false
	Barrel string `json:"barrel,omitempty" validate:"omitempty,ascii"`
	// The weapon's laser attachment
	//
	// required: false
	Laser string `json:"laser,omitempty" validate:"omitempty,ascii"`
	// The weapon's optic attachment
	//
	// required: false
	Optic string `json:"optic,omitempty" validate:"omitempty,ascii"`
	// The weapon's stock attachment.
	// N.B.: if the weapon is equipped with "No Stock," please enter that as it's stock.
	// If a stock attachment is *not* set, use an empty string.
	//
	// required: false
	// example: "No Stock"
	Stock string `json:"stock,omitempty" validate:"omitempty,ascii"`
	// The weapon's underbarrel attachment
	//
	// required: false
	Underbarrel string `json:"under_barrel,omitempty" validate:"omitempty,ascii"`
	// The weapon's magazine attachment
	//
	// required: false
	Ammo string `json:"ammo,omitempty" validate:"omitempty,ascii"`
	// The weapon's grip attachment
	//
	// required: false
	RearGrip string `json:"rear_grip,omitempty" validate:"omitempty,ascii"`
	// The weapon's perk
	//
	// required: false
	Perk string `json:"perk,omitempty" validate:"omitempty,ascii"`
	// Marks whether the weapon is a meta weapon or not
	//
	// required: true
	Meta *bool `json:"meta_weapon" validate:"required"`
}

// type Weapons is a slice of type *Weapon
type Weapons []*Weapon

// validCats is a map of all valid categories a weapon could have
var validCats map[string]bool = map[string]bool{"RANGE": true, "CLOSE": true, "CLOSE-MED": true, "SNIPER RANGED": true, "SNIPER SUPPORT": true, "SECONDARY": true}

// validGames is a map of all valid games a weapon could be from
var validGames map[string]bool = map[string]bool{"COLD WAR": true, "MODERN WARFARE": true, "VANGUARD": true}

func (w *Weapon) Validate() error {
	v := validator.New()
	return v.Struct(v)
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

func ValidGame(game string) (string, bool) {
	validGame := transformGame(game)
	_, valid := validGames[strings.ToUpper(validGame)]
	if valid {
		return validGame, true
	}
	return "", false
}

func transformGame(g string) string {
	switch {
	case strings.EqualFold(g, "coldwar"):
		return "cold war"
	case strings.EqualFold(g, "modernwarfare"):
		return "modern warfare"
	}

	return g
}

// FromJSON takes in an io.Reader, the *http.Request body,
// and unmarshals that body into a Loadout.
func (l *Weapons) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(l)
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (l *Weapons) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(l)
}

func GetStaticWeapons() Weapons {
	return StaticWeapons
}

// TODO remove as we get Google Cloud Datastore in operation
func GetMetaWeapons() Weapons {
	MetaWeapons := Weapons{}
	for _, w := range StaticWeapons {
		// have to dereference Meta field since it's a pointer
		// see: https://github.com/go-playground/validator/issues/142 &&
		// https://stackoverflow.com/questions/28817992/how-to-set-bool-pointer-to-true-in-struct-literal
		if *w.Meta == true {
			MetaWeapons = append(MetaWeapons, w)
		}
	}

	return MetaWeapons
}

// TODO remove as we get Google Cloud Datastore in operation
func GetWeaponsByGame(game string) Weapons {
	weapons := Weapons{}
	for _, w := range StaticWeapons {
		if strings.EqualFold(w.Game, game) {
			weapons = append(weapons, w)
		}
	}

	return weapons
}

// TODO remove as we get Google Cloud Datastore in operation
func GetWeaponsByName(name string) Weapons {
	weapons := Weapons{}
	for _, w := range StaticWeapons {
		if strings.EqualFold(w.WeaponName, name) {
			weapons = append(weapons, w)
		}
	}

	return weapons
}

// TODO remove as we get Google Cloud Datastore in operation
func GetWeaponsByCategory(cat string) Weapons {
	weapons := Weapons{}
	for _, w := range StaticWeapons {
		if strings.EqualFold(w.Category, cat) {
			weapons = append(weapons, w)
		}
	}

	return weapons
}

var XM4 = Weapon{
	WeaponName:  "XM4",
	Category:    "Range",
	Game:        "Cold War",
	Muzzle:      "Suppressed",
	Barrel:      "Long Barrel",
	Laser:       "",
	Optic:       "Long Sight",
	Stock:       "",
	Underbarrel: "Agent Grip",
	Ammo:        "60 round mags",
	RearGrip:    "",
	Perk:        "",
	Meta:        &[]bool{true}[0],
}

var C58 = Weapon{
	WeaponName:  "C58",
	Category:    "SniperSupport",
	Game:        "Cold War",
	Muzzle:      "Suppressed",
	Barrel:      "Long Barrel",
	Laser:       "",
	Optic:       "Long Sight",
	Stock:       "",
	Underbarrel: "Agent Grip",
	Ammo:        "45 round mags",
	RearGrip:    "",
	Perk:        "",
	// see: https://github.com/go-playground/validator/issues/142 &&
	// https://stackoverflow.com/questions/28817992/how-to-set-bool-pointer-to-true-in-struct-literal
	Meta: &[]bool{true}[0],
}

var Mac10 = Weapon{
	WeaponName:  "MAC10",
	Category:    "Close-Med",
	Muzzle:      "Suppressed",
	Barrel:      "Short Barrel",
	Laser:       "Tiger Team",
	Optic:       "Red Dot Sight",
	Stock:       "",
	Underbarrel: "Agent Grip",
	Ammo:        "60 round mags",
	RearGrip:    "",
	Perk:        "",
	// see: https://github.com/go-playground/validator/issues/142 &&
	// https://stackoverflow.com/questions/28817992/how-to-set-bool-pointer-to-true-in-struct-literal
	Meta: &[]bool{true}[0],
}

var Pistol = Weapon{
	WeaponName:  "DEAGLE",
	Category:    "Secondary",
	Game:        "Modern Warfare",
	Muzzle:      "",
	Barrel:      "Long Barrel",
	Laser:       "No Laser",
	Optic:       "Long Sight",
	Stock:       "",
	Underbarrel: "Agent Grip",
	Ammo:        "10 round mags",
	RearGrip:    "",
	Perk:        "",
	// see: https://github.com/go-playground/validator/issues/142 &&
	// https://stackoverflow.com/questions/28817992/how-to-set-bool-pointer-to-true-in-struct-literal
	Meta: &[]bool{false}[0],
}

var StaticWeapons = Weapons{&XM4, &C58, &Mac10, &Pistol}
