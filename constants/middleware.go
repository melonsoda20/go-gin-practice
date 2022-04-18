package constants

type middlewareKeys struct {
	FirebaseAppKey    string
	FirebaseClientKey string
	RedisAppKey       string
}

var MiddlewareKeysConst = middlewareKeys{
	FirebaseAppKey:    "firebase_db",
	FirebaseClientKey: "firestore_client",
	RedisAppKey:       "redis",
}
