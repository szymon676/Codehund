package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
		log.Error("create user err:", err)
		c.Redirect("/profile")
	} else {
		log.Info("successfully registered user")
		c.Redirect("/profile")
	}
	return nil
}

func (h *Handler) login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	sessionID, err := h.sm.Login(email, password)
	if err != nil {
		log.Error("err while logging in:", err)
		c.Redirect("/profile")
	}

	c.Cookie(&fiber.Cookie{
		Name:  sessionCookieName,
		Value: sessionID,
	})

	log.Info("successfully logged in")
	return c.Redirect("/profile")
}

func (h *Handler) logout(c *fiber.Ctx) error {
	sessionID := c.Cookies(sessionCookieName)

	err := h.sm.Logout(sessionID)
	if err != nil {
		log.Error("err while logging out")
		c.Redirect("/profile")
	}

	log.Info("logged out successfully")

	c.Cookie(&fiber.Cookie{
		Name:   sessionCookieName,
		Value:  "",
		MaxAge: -1,
	})

	c.Redirect("/login")
	return nil
}
