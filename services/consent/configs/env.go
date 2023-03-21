package configs

import (
	"os"
)

func EnvMongoURI() string {
	return os.Getenv("MONGO_URL")
}

func EnvMongoUsername() string {
	return os.Getenv("MONGO_USERNAME")
}

func EnvMongoPassword() string {
	return os.Getenv("MONGO_PASSWORD")
}

func EnvMongoDBName() string {
	return os.Getenv("MONGO_DB_NAME")
}
