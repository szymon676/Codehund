package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/szymon676/codehund/api"
	"github.com/szymon676/codehund/config"
	"github.com/szymon676/codehund/service"
)

func main() {
	opts := config.GetDatabaseConnectionOptions()
	app := fiber.New()
	svc := service.NewUserService(opts)
	handler := api.NewHandler(svc)

	app.Get("/", handler.RenderIndex)
	app.Get("/profile", handler.RenderProfile)
	app.Post("/register", handler.Register)

	app.Listen(":3000")
}
