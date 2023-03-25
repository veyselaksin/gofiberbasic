package config

import "os"

var (
	// MongoDBURI is the URI to connect to MongoDB
	MONGO_USER = os.Getenv("MONGO_USER")
	MONGO_PASS = os.Getenv("MONGO_PASS")
	MONGO_HOST = os.Getenv("MONGO_HOST")
	MONGO_PORT = os.Getenv("MONGO_PORT")
	MONGO_DB   = os.Getenv("MONGO_DB")
	MONGO_URI  = "mongodb://" + MONGO_USER + ":" + MONGO_PASS + "@" + MONGO_HOST + ":" + MONGO_PORT + "/" + MONGO_DB + "?authSource=admin"
)
