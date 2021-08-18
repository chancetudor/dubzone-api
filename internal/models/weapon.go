package models

import (
	p "go.mongodb.org/mongo-driver/bson/primitive"
)

// Weapon represents a Warzone weapon.
type Weapon struct {
	ID             p.ObjectID    `json:"_id,omitempty" bson:"_id,omitempty"`
	WeaponName     string        `json:"weapon_name,omitempty" bson:"weapon_name,omitempty"`
	GameFrom       string        `json:"game_from,omitempty" bson:"game_from,omitempty"`
	RPM            int           `json:"rpm,omitempty" bson:"rpm,omitempty"`
	BulletVelocity int           `json:"bullet_velocity,omitempty" bson:"bullet_velocity,omitempty"`
	DamageProfile  DamageProfile `json:"damage_profile,omitempty" bson:"damage_profile,omitempty"`
}
