package api

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) withAuth(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionID := c.Cookies(sessionCookieName)
		if sessionID == "" {
			return c.Redirect("/login")
		}
		return next(c)
	}
}

func (h *Handler) InitRoutes() {
	app := fiber.New()

	app.Get("/", h.renderIndex)
	app.Get("/profile/:username", h.renderUserByUsername)
	app.Get("/login", h.renderLogin)
	app.Get("/register", h.renderRegister)

	app.Post("/register", h.register)
	app.Post("/login", h.login)
	app.Post("/logout", h.logout)

	app.Post("/posts", h.createPost)

	app.Delete("/posts/:postid", h.deletePost)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":3000"
	}
	app.Listen(port)
}
