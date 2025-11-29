package util

import "github.com/gofiber/fiber/v2"

func Error(c *fiber.Ctx, code int, msg string) error {
	return c.Status(code).JSON(fiber.Map{"error": msg})
}

func Success(c *fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{"data": data})
}
