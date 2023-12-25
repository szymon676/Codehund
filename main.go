package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/szymon676/codehund/api"
)

func main() {
	app := fiber.New()
	handler := api.Handler{}

	app.Get("/", handler.RenderIndex)
	app.Post("/register", handler.Register)

	// app.Post("/register")
	// app.Post("/login")
	// app.Get("/user/:id")

	// app.Get("/posts")
	// app.Post("/post")
	// app.Delete("/post/:id")

	// app.Post("/comment")
	// app.Get("/comments/:id")
	// app.Delete("/comment/:id")

	app.Listen(":3000")
}
