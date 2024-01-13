package api

import (
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
		return c.SendString("Could not create post. Please try again")
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
		return c.SendString("Error, did not get post id properly.")
	}
	err = h.svc.DeletePost(convpostID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error deleting post, please try again.")
	}
	return c.SendString("post deleted succesfuly")
}
