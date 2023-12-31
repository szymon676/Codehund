package api

import "github.com/gofiber/fiber/v2"

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
	app.Get("/profile", h.withAuth(h.renderProfile))
	app.Get("/login", h.renderLogin)
	app.Get("/register", h.renderRegister)

	app.Post("/register", h.register)
	app.Post("/login", h.login)
	app.Post("/logout", h.logout)

	app.Post("/posts", h.createPost)

	app.Listen(":3000")
}
