package main

import (
	"github.com/gofiber/fiber/v2"
)

//c.OriginalURL() 返回原始的请求数据

func main() {

	app := fiber.New()

	// GET http://example.com/search?q=something
	app.Get("/*", func(c *fiber.Ctx) error {
		s := c.OriginalURL() // "/search?q=something"

		return c.SendString(s)
		// ...
	})

	app.Listen(":8080")
}
