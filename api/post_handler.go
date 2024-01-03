package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/szymon676/codehund/types"
)

func (h *Handler) createPost(c *fiber.Ctx) error {
	sessionID := c.Cookies(sessionCookieName)
	content := c.FormValue("content")
	author, err := h.sm.GetSession(sessionID)
	if err != nil {
		return c.Redirect("/login")
	}
	post := &types.Post{
		Content: content,
		Author:  author.Username,
		Date:    time.Now(),
	}
	err = h.svc.CreatePost(post)
	if err != nil {
		return err
	}
	return c.Redirect("/")
}
