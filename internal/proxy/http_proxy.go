package proxy

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func Forward(target string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		url := target + c.Params("*")
		return proxy.Do(c, url)
	}
}
