package route

import (
	"github.com/ChinawatDc/001-fiber-intro-golang-fiber/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	h := handler.NewHealthHandler()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/health", h.Health)

	v1.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name", "world")
		return c.JSON(fiber.Map{
			"message": "hello " + name,
			"method":  c.Method(),
			"path":    c.Path(),
			"rid":     c.Locals("request_id"),
		})
	})
}
