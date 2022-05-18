package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	//一个fiber的应用程序
	app := fiber.New()

	//请求根路径
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//启动应用
	app.Listen("127.0.0.1:8080")
}
