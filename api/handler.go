package api

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/szymon676/codehund/service"
	"github.com/szymon676/codehund/types"
	"github.com/szymon676/codehund/views/index"
	"github.com/szymon676/codehund/views/profile"
)

type Handler struct {
	svc service.UserServicer
}

func NewHandler(svc service.UserServicer) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (*Handler) RenderIndex(c *fiber.Ctx) error {
	return render(c, index.Show())
}

func (*Handler) RenderProfile(c *fiber.Ctx) error {
	userstate := &types.Userstate{
		Loggedin: true,
	}
	return render(c, profile.Show(userstate))
}

func (h *Handler) Register(c *fiber.Ctx) error {
	name := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	user := &types.User{
		Username: name,
		Password: password,
		Email:    email,
	}
	err := h.svc.CreateUser(user)
	if err != nil {
		log.Info(err)
		c.JSON(fiber.Map{"error": err.Error()})
	} else {
		return c.SendString("sucessfuly registered user")
	}
	return nil
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
