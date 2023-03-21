package configs

import (
	"fmt"
	"os"
)

func EnvMongoURI() string {
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_DB_HOST")
	port := os.Getenv("MONGO_PORT")

	return fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
}

func EnvMongoDBName() string {
	return os.Getenv("MONGO_DB_NAME")
}
