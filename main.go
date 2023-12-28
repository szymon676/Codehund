package main

import (
	"github.com/szymon676/codehund/api"
	"github.com/szymon676/codehund/auth"
	"github.com/szymon676/codehund/config"
	"github.com/szymon676/codehund/db"
	"github.com/szymon676/codehund/service"
)

func main() {
	psqlopts := config.GetDatabaseConnectionOptions()
	rdbopts := config.GetRedisConnOptions()

	psqldb := db.NewPostgresDatabase(psqlopts)
	rdb := db.NewRedisClient(rdbopts)

	svc := service.NewUserService(psqldb)
	session := auth.NewSessionManager(rdb, psqldb)
	handler := api.NewHandler(svc, session)
	handler.InitRoutes()
}
