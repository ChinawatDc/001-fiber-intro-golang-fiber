package middleware

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
)

const headerRequestID = "X-Request-Id"

func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		rid := c.Get(headerRequestID)
		if rid == "" {
			rid = newRequestID()
		}
		c.Locals("request_id", rid)

		c.Set(headerRequestID, rid)

		return c.Next()
	}
}

func newRequestID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
