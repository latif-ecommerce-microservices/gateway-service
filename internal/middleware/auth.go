package middleware

import "github.com/gofiber/fiber/v2"

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "missing Authorization token",
			})
		}
		return c.Next()
	}
}
