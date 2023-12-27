package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/szymon676/codehund/types"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env")
	}
}

func GetDatabaseConnectionOptions() *types.PostgresConnectionOptions {
	user := os.Getenv("SQL_USER")
	databaseName := os.Getenv("SQL_DATABASENAME")
	port := os.Getenv("SQL_PORT")
	password := os.Getenv("SQL_PASSWORD")

	return &types.PostgresConnectionOptions{
		User:         user,
		DatabaseName: databaseName,
		Port:         port,
		Password:     password,
	}
}

func GetRedisConnOptions() *types.RedisConnectionOptions {
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")
	return &types.RedisConnectionOptions{
		Port:     port,
		Password: password,
	}
}
