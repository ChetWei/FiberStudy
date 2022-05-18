package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//存储请求范围内变量的方法，试用与将特定数据传给下一个中间件

func main() {
	app := fiber.New()

	//任意请求都会进入
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("app use")
		//设置一个值，传递给下一个路由
		c.Locals("user", "admin")
		return c.Next()
	})

	app.Get("/admin", func(c *fiber.Ctx) error {
		//获取local变量
		if c.Locals("user") == "admin" {
			return c.Status(fiber.StatusOK).SendString("Welcome, admin!")
		}
		return c.SendStatus(fiber.StatusForbidden)

	})

	app.Listen(":8080")
}
