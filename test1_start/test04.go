package main

import "github.com/gofiber/fiber/v2"

func main() {
	//一个fiber的实例
	app := fiber.New()

	//获取 /api后面的所有值
	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
	})

	//启动应用
	app.Listen("127.0.0.1:8080")

}
