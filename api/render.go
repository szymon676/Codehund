package api

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/szymon676/codehund/auth"
	"github.com/szymon676/codehund/views/pages/index"
	"github.com/szymon676/codehund/views/pages/login"
	"github.com/szymon676/codehund/views/pages/profile"
	"github.com/szymon676/codehund/views/pages/register"
)

func (h *Handler) renderIndex(c *fiber.Ctx) error {
	sessionID := c.Cookies(sessionCookieName)
	user, err := h.sm.GetSession(sessionID)
	if err != nil {
		user = new(auth.UserSession)
		user.Username = ""
	}
	posts, err := h.svc.GetPosts()
	if err != nil {
		return err
	}
	return render(c, index.Show(posts, user.Username))
}

func (h *Handler) renderProfile(c *fiber.Ctx) error {
	sessionID := c.Cookies(sessionCookieName)
	user, err := h.sm.GetSession(sessionID)
	if err != nil {
		return c.Redirect("/login")
	}
	return render(c, profile.Show(user.Username))
}

func (h *Handler) renderRegister(c *fiber.Ctx) error {
	return render(c, register.Show())
}

func (h *Handler) renderLogin(c *fiber.Ctx) error {
	return render(c, login.Show())
}

func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
