package config

import (
	"os"

	"github.com/szymon676/codehund/types"
)

func GetDatabaseConnectionOptions() *types.ConnectionOptions {
	user := os.Getenv("SQL_USER")
	databaseName := os.Getenv("SQL_DATABASENAME")
	port := os.Getenv("SQL_PORT")
	password := os.Getenv("SQL_PASSWORD")

	return &types.ConnectionOptions{
		User:         user,
		DatabaseName: databaseName,
		Port:         port,
		Password:     password,
	}
}
