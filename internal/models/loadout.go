package models

// Loadout represents a Warzone loadout.
type Loadout struct {
	Weapon      string `json:"weapon,omitempty" bson:"weapon,omitempty"`
	Category    string `json:"category,omitempty" bson:"category,omitempty"`
	Muzzle      string `json:"muzzle,omitempty" bson:"muzzle,omitempty"`
	Barrel      string `json:"barrel,omitempty" bson:"barrel,omitempty"`
	Laser       string `json:"laser,omitempty" bson:"laser,omitempty"`
	Optic       string `json:"optic,omitempty" bson:"optic,omitempty"`
	Stock       string `json:"stock,omitempty" bson:"stock,omitempty"`
	Underbarrel string `json:"under_barrel,omitempty" bson:"under_barrel,omitempty"`
	Ammo        string `json:"ammo,omitempty" bson:"ammo,omitempty"`
	RearGrip    string `json:"rear_grip,omitempty" bson:"rear_grip,omitempty"`
	Perk        string `json:"perk,omitempty" bson:"perk,omitempty"`
}
