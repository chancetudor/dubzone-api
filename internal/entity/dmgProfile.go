package entity

// DamageProfile models the close-, mid-, and far-range time and shots to kill
type DamageProfile struct {
	CloseRange struct {
		// just max distance here because min distance = 0
		MaxDistance 	int 	`json:"max_distance,omitempty" bson:"max_distance,omitempty"`
		MinShotsToKill 	int		`json:"min_stk,omitempty" bson:"min_stk,omitempty"`
		MaxShotsToKill 	int 	`json:"max_stk,omitempty" bson:"max_stk,omitempty"`
		MinTimeToKill 	int 	`json:"min_ttk,omitempty" bson:"min_ttk,omitempty"`
		MaxTimeToKill 	int 	`json:"max_ttk,omitempty" bson:"max_ttk,omitempty"`
	} 							`json:"close_range,omitempty" bson:"close_range,omitempty"`
	MidRange struct {
		// min distance and max distance to define a true "mid" range
		MinDistance 	int 	`json:"min_distance,omitempty" bson:"min_distance,omitempty"`
		MaxDistance 	int 	`json:"max_distance,omitempty" bson:"max_distance,omitempty"`
		MinShotsToKill 	int		`json:"min_stk,omitempty" bson:"min_stk,omitempty"`
		MaxShotsToKill 	int 	`json:"max_stk,omitempty" bson:"max_stk,omitempty"`
		MinTimeToKill 	int 	`json:"min_ttk,omitempty" bson:"min_ttk,omitempty"`
		MaxTimeToKill 	int 	`json:"max_ttk,omitempty" bson:"max_tkk,omitempty"`
	} 							`json:"mid_range,omitempty" bson:"mid_range,omitempty"`
	FarRange struct {
		// just min distance here because max distance = infinity (theoretically)
		MinDistance		int		`json:"min_distance,omitempty" bson:"min_distance,omitempty"`
		MinShotsToKill 	int		`json:"min_stk,omitempty" bson:"min_stk,omitempty"`
		MaxShotsToKill 	int 	`json:"max_stk,omitempty" bson:"max_stk,omitempty"`
		MinTimeToKill 	int 	`json:"min_ttk,omitempty" bson:"min_ttk,omitempty"`
		MaxTimeToKill 	int 	`json:"max_ttk,omitempty" bson:"max_ttk,omitempty"`
	} 							`json:"far_range,omitempty" bson:"far_range,omitempty"`
}
