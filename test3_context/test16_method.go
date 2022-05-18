package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//获取客户端的请求消息

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		//获取客户端的请求方式
		fmt.Println(c.Method()) // "POST"
		//修改请求方式
		c.Method("GET")
		fmt.Println(c.Method()) // GET

		// ...
		return nil
	})

	app.Listen(":8080")
}
