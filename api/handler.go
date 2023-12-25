package api

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/szymon676/codehund/views/index"
)

type Handler struct{}

func (*Handler) RenderIndex(c *fiber.Ctx) error {
	return render(c, index.Show())
}

func (*Handler) Register(c *fiber.Ctx) error {
	var req any
	err := c.BodyParser(&req)
	if err != nil {
		return c.SendString("could not read req body")
	}
	return c.SendString("sucessfuly registered user")
}

func (*Handler) Login(c *fiber.Ctx) error {
	return nil
}

func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
