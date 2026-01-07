package main

import (
	"log"

	"github.com/ChinawatDc/001-fiber-intro-golang-fiber/internal/route"
	"github.com/ChinawatDc/001-fiber-intro-golang-fiber/internal/server"
)

func main() {
	app := server.NewFiber()

	route.RegisterRoutes(app)

	log.Println("🚀 Fiber API is running on http://localhost:8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
