package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/szymon676/codehund/types"
)

func (h *Handler) createPost(c *fiber.Ctx) error {
	sessionID := c.Cookies(sessionCookieName)
	title := c.FormValue("title")
	content := c.FormValue("content")
	author, err := h.sm.GetSession(sessionID)
	if err != nil {
		return c.Redirect("/login")
	}
	post := &types.Post{
		Title:   title,
		Content: content,
		Author:  author.Username,
	}
	err = h.svc.CreatePost(post)
	if err != nil {
		return c.Redirect("/")
	}
	return c.Redirect("/")
}
