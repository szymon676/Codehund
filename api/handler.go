package api

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/szymon676/codehund/auth"
	"github.com/szymon676/codehund/service"
	"github.com/szymon676/codehund/types"
	"github.com/szymon676/codehund/views/pages/index"
	"github.com/szymon676/codehund/views/pages/login"
	"github.com/szymon676/codehund/views/pages/profile"
)

type Handler struct {
	svc       service.UserServicer
	sm        *auth.SessionManager
	sessionid string
}

func NewHandler(svc service.UserServicer, sm *auth.SessionManager) *Handler {
	return &Handler{
		svc:       svc,
		sm:        sm,
		sessionid: "",
	}
}

func (h *Handler) InitRoutes() {
	app := fiber.New()
	app.Get("/", h.RenderIndex)
	app.Get("/profile", h.withAuth(h.RenderProfile))
	app.Get("/login", h.RenderLogin)
	app.Post("/register", h.Register)
	app.Post("/login", h.Login)
	app.Post("/logout", h.Logout)

	app.Listen(":3000")
}

func (h *Handler) withAuth(next fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if h.sessionid == "" {
			return c.Redirect("/login")
		}
		return next(c)
	}
}

func (h *Handler) RenderIndex(c *fiber.Ctx) error {
	return render(c, index.Show())
}

func (h *Handler) RenderProfile(c *fiber.Ctx) error {
	user, err := h.sm.GetSession(h.sessionid)
	if err != nil {
		c.Redirect("/login")
	}
	return render(c, profile.Show(user.Username))
}

func (h *Handler) RenderLogin(c *fiber.Ctx) error {
	return render(c, login.Show())
}

func (h *Handler) Register(c *fiber.Ctx) error {
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
		log.Info("create user err:", err)
		c.Redirect("/profile")
	} else {
		log.Info("successfully registered user")
		c.Redirect("/profile")
	}

	return nil
}

func (h *Handler) Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	sessionid, err := h.sm.Login(email, password)
	if err != nil {
		log.Info("err while logging in:", err)
		c.Redirect("/profile")
	}
	h.sessionid = sessionid
	log.Info("successfully logged in")
	return c.Redirect("/profile")
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	h.sm.Logout(h.sessionid)
	h.sessionid = ""
	c.Redirect("/profile")
	return nil
}

func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
