package main

import (
	"github.com/gofiber/fiber/v2"
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

	app := fiber.New()
	app.Get("/", handler.RenderIndex)
	app.Get("/profile", handler.RenderProfile)
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
	app.Post("/logout", handler.Logout)

	app.Listen(":3000")
}
