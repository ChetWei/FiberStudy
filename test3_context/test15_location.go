package main

import (
	"github.com/gofiber/fiber/v2"
)

//设置响应头中  Location: value

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Location("/test")
		return nil
	})

	app.Listen(":8080")
}
