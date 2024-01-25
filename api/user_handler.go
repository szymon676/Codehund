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
		c.Redirect("/login")
	} else {

		c.Cookie(&fiber.Cookie{
			Name:  sessionCookieName,
			Value: sessionID,
		})

		return c.Redirect("/")
	}
	return nil
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

	return c.SendString("logged out")
}

// todo
func (h *Handler) follow(c *fiber.Ctx) error {
	sessionID := c.Cookies(sessionCookieName)
	follower, err := h.sm.GetSession(sessionID)
	if err != nil {
		return c.SendString("Could not follow user")
	}

	followeeUsername := c.FormValue("followee")
	followee, err := h.svc.GetUserByUsername(followeeUsername)
	if err != nil {
		return c.SendString("Could not follow user")
	}

	err = h.svc.FollowUser(follower.ID, followee.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("could not followe user.")
	}

	return c.SendString("followed user")
}

func (h *Handler) unfollow(c *fiber.Ctx) error {
	return nil
}
