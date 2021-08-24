package models

import (
	p "go.mongodb.org/mongo-driver/bson/primitive"
)

// Loadout represents a Warzone loadout,
// complete with a primary weapon, secondary weapon, three perks, and lethal and tactical equipment
type Loadout struct {
	ID        p.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Primary   Weapon     `json:"primary,omitempty" bson:"primary,omitempty"`
	Secondary Weapon     `json:"secondary,omitempty" bson:"secondary,omitempty"`
	PerkOne   string     `json:"perk_one,omitempty" bson:"perk_one,omitempty"`
	PerkTwo   string     `json:"perk_two,omitempty" bson:"perk_two,omitempty"`
	PerkThree string     `json:"perk_three,omitempty" bson:"perk_three,omitempty"`
	Lethal    string     `json:"lethal,omitempty" bson:"lethal,omitempty"`
	Tactical  string     `json:"tactical,omitempty" bson:"tactical,omitempty"`
	Meta      bool       `json:"meta_loadout,omitempty" bson:"meta_loadout,omitempty"`
}
