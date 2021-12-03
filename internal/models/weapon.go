package models

import "github.com/go-playground/validator/v10"

// Weapon represents a Warzone weapon, complete with a category and all recommended attachments.
type Weapon struct {
	WeaponName  string `json:"weapon_name,omitempty" validate:"required, alpha, max=15"`
	Category    string `json:"category,omitempty" validate:"required, alpha"`
	Muzzle      string `json:"muzzle,omitempty" validate:"required, alpha"`
	Barrel      string `json:"barrel,omitempty" validate:"required, alpha"`
	Laser       string `json:"laser,omitempty" validate:"required, alpha"`
	Optic       string `json:"optic,omitempty" validate:"required, alpha"`
	Stock       string `json:"stock,omitempty" validate:"required, alpha"`
	Underbarrel string `json:"under_barrel,omitempty" validate:"required, alpha"`
	Ammo        string `json:"ammo,omitempty" validate:"required, alpha"`
	RearGrip    string `json:"rear_grip,omitempty" validate:"required, alpha"`
	Perk        string `json:"perk,omitempty" validate:"required, alpha"`
	Meta        bool   `json:"meta_weapon,omitempty" validate:"required, boolean"`
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
	Meta:        true,
}

var C58 = Weapon{
	WeaponName:  "C58",
	Category:    "Range",
	Muzzle:      "Suppressed",
	Barrel:      "Long Barrel",
	Laser:       "",
	Optic:       "Long Sight",
	Stock:       "",
	Underbarrel: "Agent Grip",
	Ammo:        "45 round mags",
	RearGrip:    "",
	Perk:        "",
	Meta:        true,
}

var Mac10 = Weapon{
	WeaponName:  "MAC 10",
	Category:    "Close Range",
	Muzzle:      "Suppressed",
	Barrel:      "Short Barrel",
	Laser:       "Tiger Team",
	Optic:       "Red Dot Sight",
	Stock:       "",
	Underbarrel: "Agent Grip",
	Ammo:        "60 round mags",
	RearGrip:    "",
	Perk:        "",
	Meta:        true,
}

var Pistol = Weapon{
	WeaponName:  "Colt .45",
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
	Meta:        false,
}

var StaticWeapons = Weapons{XM4, C58, Mac10, Pistol}
