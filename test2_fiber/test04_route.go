package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//路由分组
func main() {
	app := fiber.New()

	//匹配请求以/api
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("/api 进入下一个路由")
		return c.Next() //进入到下一个路由
	})

	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("/api")
		return c.SendString("/api")
	})

	//简单GET 处理
	app.Get("/api/list", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})

	// 简单POST处理
	app.Post("/api/register", func(c *fiber.Ctx) error {
		return c.SendString("I'm a POST request!")
	})

	app.Listen(":8080")
}
