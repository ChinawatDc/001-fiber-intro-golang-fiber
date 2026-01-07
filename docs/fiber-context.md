# Fiber Context (fiber.Ctx) คืออะไร

Fiber ใช้ `*fiber.Ctx` เป็นตัวแทน "request + response" ของคำขอหนึ่งครั้ง

## สิ่งที่ใช้บ่อย
- Method / Path
  - c.Method()
  - c.Path()

- Query string
  - c.Query("key")
  - c.Query("key", "default")

- Path params
  - c.Params("id")

- Headers
  - c.Get("Authorization")
  - c.Set("X-Request-Id", "abc")

- Body (JSON)
  - var req MyReq
  - c.BodyParser(&req)

- Response
  - c.Status(200).JSON(...)
  - c.SendString("ok")

- Per-request storage
  - c.Locals("request_id", "xxx")
  - c.Locals("request_id")
