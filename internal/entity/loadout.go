package entity

// Loadout represents a Warzone loadout.
type Loadout struct {
	Category	string		`json:"category,omitempty" bson:"category,omitempty"`
	Muzzle		string		`json:"muzzle,omitempty" bson:"muzzle,omitempty"`
	Barrel		string		`json:"barrel,omitempty" bson:"barrel,omitempty"`
	Laser 		string		`json:"laser,omitempty" bson:"laser,omitempty"`
	Optic 		string		`json:"optic,omitempty" bson:"optic,omitempty"`
	Stock		string		`json:"stock,omitempty" bson:"stock,omitempty"`
	Underbarrel string		`json:"underbarrel,omitempty" bson:"underbarrel,omitempty"`
	Ammo 		string		`json:"ammo,omitempty" bson:"ammo,omitempty"`
	RearGrip	string		`json:"reargrip,omitempty" bson:"reargrip,omitempty"`
	Perk		string		`json:"perk,omitempty" bson:"perk,omitempty"`
}
