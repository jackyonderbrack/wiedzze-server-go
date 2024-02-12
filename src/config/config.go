package config

import "os"

var (
	MongoURL = os.Getenv("MONGO_URL")
)