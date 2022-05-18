package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//获取域名后面的path内容，不包括 ? 后面的内容

func main() {
	app := fiber.New()

	// GET http://example.com/users?sort=desc

	app.Get("/users", func(c *fiber.Ctx) error {

		fmt.Println(c.Path())
		//重写
		c.Path("/john")
		c.Path() // "/john"

		return nil
	})

	app.Listen(":8080")
}
