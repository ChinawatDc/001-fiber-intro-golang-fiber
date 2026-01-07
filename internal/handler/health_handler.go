package handler

import "github.com/gofiber/fiber/v2"

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// GET /health
func (h *HealthHandler) Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"rid":    c.Locals("request_id"),
	})
}
