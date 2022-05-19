package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world！proxy2")
	})

	app.Listen(":8081")
}
