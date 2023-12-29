package api

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/szymon676/codehund/views/pages/index"
	"github.com/szymon676/codehund/views/pages/login"
	"github.com/szymon676/codehund/views/pages/profile"
	"github.com/szymon676/codehund/views/pages/register"
)

func (h *Handler) renderIndex(c *fiber.Ctx) error {
	return render(c, index.Show())
}

func (h *Handler) renderProfile(c *fiber.Ctx) error {
	if h.sessionid == "" {
		return c.Redirect("/login")
	}
	user, err := h.sm.GetSession(h.sessionid)
	if err != nil {
		return c.Redirect("/login")
	}
	return render(c, profile.Show(user.Username))
}

func (h *Handler) renderRegister(c *fiber.Ctx) error {
	return render(c, register.Show())
}

func (h *Handler) renderLogin(c *fiber.Ctx) error {
	if h.sessionid != "" {
		c.Redirect("/profile")
	}
	return render(c, login.Show())
}

func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
