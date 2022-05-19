package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

//func handler(c *fiber.Ctx) error {
//	return c.JSON("{'hello':'world!'}")
//}

func main() {
	app := fiber.New()

	app.Get("/john/:age", handler)
	app.Post("/register", handler)

	//获取路由器的堆栈信息
	data, _ := json.MarshalIndent(app.Stack(), "", "  ")

	app.Get("/info", func(ctx *fiber.Ctx) error {
		return ctx.SendString(string(data))
	})

	app.Listen(":3000")
}
