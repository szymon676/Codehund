package api

import (
	"log"
	"strconv"
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
		return c.Redirect("/")
	}
	return c.Redirect("/")
}

func (h *Handler) deletePost(c *fiber.Ctx) error {
	sessionID := c.Cookies(sessionCookieName)
	_, err := h.sm.GetSession(sessionID)
	if err != nil {
		return c.Redirect("/login")
	}
	postID := c.Params("postid")
	convpostID, err := strconv.Atoi(postID)
	if err != nil {
		return c.Redirect("/")
	}
	err = h.svc.DeletePost(convpostID)
	if err != nil {
		log.Println(err)
	}
	return c.Redirect("/")
}
