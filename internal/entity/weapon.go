package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Weapon represents an Warzone weapon.
type Weapon struct {
	ID				primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	WeaponName		string				`json:"weaponname,omitempty" bson:"weaponname,omitempty"`
	GameFrom		string 				`json:"gamefrom,omitempty" bson:"gamefrom,omitempty"`
	RPM 			int 				`json:"rpm,omitempty" bson:"rpm,omitempty"`
	BulletVelocity	int					`json:"bulletvelocity,omitempty" bson:"bulletvelocity,omitempty"`
	DamageProfile 	DamageProfile		`json:"damageprofile,omitempty" bson:"damageprofile,omitempty"`
	Loadouts		[]Loadout			`json:"loadouts,omitempty" bson:"loadouts,omitempty"`
}
