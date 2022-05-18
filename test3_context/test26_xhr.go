package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// X-Requested-With: XMLHttpRequest

	app.Get("/", func(c *fiber.Ctx) error {
		r := c.XHR() // 判断请求是否为xhr
		print(r)
		// ...
		return nil
	})

	app.Listen(":8080")
}
