package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//c.Next()方法，
//当调用Next时，他会执行堆栈中与当前路由匹配的下一个方法

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("1st route!")
		return c.Next()
	})

	app.Get("*", func(c *fiber.Ctx) error {
		fmt.Println("2nd route!")
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("3rd route!")
		return c.SendString("Hello, World!")
	})

	app.Listen(":8080")
}
