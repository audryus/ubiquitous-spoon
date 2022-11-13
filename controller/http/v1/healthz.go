package v1

import "github.com/gofiber/fiber/v2"

func (h *Handler) healthz(v1 fiber.Router) {
	v1.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
}
