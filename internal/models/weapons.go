package models

import "github.com/go-playground/validator/v10"

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

type Weapons []Weapon

func (w *Weapon) Validate() error {
	v := validator.New()
	return v.Struct(v)
}

func GetWeapons() Weapons {
	return StaticWeapons
}

var XM4 = Weapon{
	WeaponName:  "XM4",
	Category:    "Range",
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
	Category:    "Sniper Support",
	Muzzle:      "Suppressed",
	Barrel:      "Long Barrel",
	Laser:       "",
	Optic:       "Long Sight",
	Stock:       "",
	Underbarrel: "Agent Grip",
	Ammo:        "45 round mags",
	RearGrip:    "",
	Perk:        "",
	Meta:        &[]bool{true}[0],
}

var Mac10 = Weapon{
	WeaponName:  "MAC 10",
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
	Meta:        &[]bool{true}[0],
}

var Pistol = Weapon{
	WeaponName:  "COLT .45",
	Category:    "Secondary",
	Muzzle:      "",
	Barrel:      "Long Barrel",
	Laser:       "No Laser",
	Optic:       "Long Sight",
	Stock:       "",
	Underbarrel: "Agent Grip",
	Ammo:        "10 round mags",
	RearGrip:    "",
	Perk:        "",
	Meta:        &[]bool{false}[0],
}

var StaticWeapons = Weapons{XM4, C58, Mac10, Pistol}
