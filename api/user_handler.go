package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/szymon676/codehund/types"
)

const sessionCookieName = "sessionID"

func (h *Handler) register(c *fiber.Ctx) error {
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
		c.Status(fiber.StatusInternalServerError).SendString("Could not register user")
	}

	return c.Redirect("/login")
}

func (h *Handler) login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	sessionID, err := h.sm.Login(email, password)
	if err != nil {
		c.SendString("error while logging in.")
	}

	c.Cookie(&fiber.Cookie{
		Name:  sessionCookieName,
		Value: sessionID,
	})

	return c.Redirect("/profile")
}

func (h *Handler) logout(c *fiber.Ctx) error {
	sessionID := c.Cookies(sessionCookieName)

	err := h.sm.Logout(sessionID)
	if err != nil {
		c.SendString("Error while logging out")
	}

	c.Cookie(&fiber.Cookie{
		Name:   sessionCookieName,
		Value:  "",
		MaxAge: -1,
	})

	return c.Redirect("/login")
}
