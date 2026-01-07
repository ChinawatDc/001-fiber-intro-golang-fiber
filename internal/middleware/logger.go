package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger() fiber.Handler {
	return fiberLogger.New(fiberLogger.Config{
		TimeFormat: time.RFC3339,
		Format: fmt.Sprintf(
			`{"time":"${time}","rid":"${locals:request_id}","ip":"${ip}","method":"${method}","path":"${path}","status":${status},"latency":"${latency}"}%s`,
			"\n",
		),
	})
}
