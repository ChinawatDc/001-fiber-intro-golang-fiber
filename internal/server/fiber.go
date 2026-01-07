package server

import (
	"github.com/ChinawatDc/001-fiber-intro-golang-fiber/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func NewFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "001-fiber-intro",
	})

	app.Use(middleware.RequestID())
	app.Use(middleware.Logger())

	return app
}
