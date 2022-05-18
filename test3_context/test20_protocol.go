package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//获取http 或 https协议

func main() {

	app := fiber.New()

	// GET http://example.com

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c.Protocol()) // "http"

		return nil
	})

	app.Listen(":8080")
}
