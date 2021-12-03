package auth

// type MongoAuth struct {
// 	URI      string
// 	Database string
// 	LoadoutCollection string
// 	WeaponsCollection string
// }
//
// // NewAuth Creates a new MongoAuth struct, containing the username and password necessary to connect to a MongoDB client
// // along with the necessary collection names
// func NewAuth() *MongoAuth {
// 	logger.Info("Retrieving auth info", "auth.NewAuth()")
// 	err := godotenv.Load("./config/dev.env")
// 	if err != nil {
// 		logger.Fatal(err, "Loading .env file", "auth.NewAuth()")
// 	}
//
// 	return &MongoAuth{
// 		URI:      os.Getenv("MONGO_URI"),
// 		Database: os.Getenv("DATABASE"),
// 		LoadoutCollection: os.Getenv("LOADOUTSCOLLECTIONNAME"),
// 		WeaponsCollection: os.Getenv("WEAPONSCOLLECTIONNAME"),
// 	}
// }
