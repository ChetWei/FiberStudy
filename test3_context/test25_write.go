package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Write([]byte("Hello, World!")) // => "Hello, World!"

		fmt.Fprintf(c, "%s\n", "Hello, World!") // "Hello, World!Hello, World!"
		return nil
	})

	//接受字符串 参数
	app.Get("/1", func(c *fiber.Ctx) error {
		world := "World!"
		c.Writef("Hello, %s", world) // => "Hello, World!"

		fmt.Fprintf(c, "%s\n", "Hello, World!") // "Hello, World!Hello, World!"
		return nil
	})

	//写入string
	app.Get("/2", func(c *fiber.Ctx) error {

		c.WriteString("Hello, World!") // => "Hello, World!"

		fmt.Fprintf(c, "%s\n", "Hello, World!") // "Hello, World!Hello, World!"
		return nil
	})

	app.Listen(":8080")
}
