package entity

// DamageProfile models the close and far-range time and shots to kill
type DamageProfile struct {
	CloseRange struct {
		// just max distance here because min distance = 0
		MaxDistance 	int 	`json:"maxdistance,omitempty" bson:"maxdistance,omitempty"`
		MinShotsToKill 	int		`json:"minstk,omitempty" bson:"minstk,omitempty"`
		MaxShotsToKill 	int 	`json:"maxstk,omitempty" bson:"maxstk,omitempty"`
		MinTimeToKill 	int 	`json:"minttk,omitempty" bson:"minttk,omitempty"`
		MaxTimeToKill 	int 	`json:"maxttk,omitempty" bson:"maxttk,omitempty"`
	} `json:"closerange,omitempty" bson:"closerange,omitempty"`
	MidRange struct {
		MinDistance 	int 	`json:"mindistancemid,omitempty" bson:"mindistancemid,omitempty"`
		MaxDistance 	int 	`json:"maxdistancemid,omitempty" bson:"maxdistancemid,omitempty"`
		MinShotsToKill 	int		`json:"minstkmid,omitempty" bson:"minstkmid,omitempty"`
		MaxShotsToKill 	int 	`json:"maxstkmid,omitempty" bson:"maxstkmid,omitempty"`
		MinTimeToKill 	int 	`json:"minttkmid,omitempty" bson:"minttkmid,omitempty"`
		MaxTimeToKill 	int 	`json:"maxttkmid,omitempty" bson:"maxttkmid,omitempty"`
	} `json:"midrange,omitempty" bson:"midrange,omitempty"`
	FarRange struct {
		// just min distance here because max distance = infinity (theoretically)
		MinDistance		int		`json:"mindistancefar,omitempty" bson:"mindistancefar,omitempty"`
		MinShotsToKill 	int		`json:"minstkfar,omitempty" bson:"minstkfar,omitempty"`
		MaxShotsToKill 	int 	`json:"maxstkfar,omitempty" bson:"maxstkfar,omitempty"`
		MinTimeToKill 	int 	`json:"minttkfar,omitempty" bson:"minttkfar,omitempty"`
		MaxTimeToKill 	int 	`json:"maxttkfar,omitempty" bson:"maxttkfar,omitempty"`
	} `json:"farrange,omitempty" bson:"farrange,omitempty"`
}
