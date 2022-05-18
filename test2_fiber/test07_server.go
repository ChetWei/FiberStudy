package main

import "github.com/gofiber/fiber/v2"

//设置 fasthttp server 信息

func main() {
	app := fiber.New()
	app.Server().MaxConnsPerIP = 1 //每个ip最多的连接数量

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("/hello方法")
	})

	app.Listen(":8080")

}
