package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{
		AppName: "test",
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		//获取实例的config的信息
		return ctx.SendString(app.Config().AppName)
	})

	app.Listen(":8080")
}
