package controller

import (
	v1 "github.com/audryus/ubiquitous-spoon/controller/http/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	h.api(app)
	return app
}

func (h *Handler) api(app *fiber.App) {
	v1 := v1.NewHandler()
	api := app.Group("/api")
	{
		v1.Init(api)
	}
}
