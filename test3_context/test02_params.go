package main

import "github.com/gofiber/fiber/v2"

//params参数获取

func main() {
	app := fiber.New()

	// GET http://example.com/user/fenny
	app.Get("/user/:name", func(c *fiber.Ctx) error {
		c.AllParams() // "{"name": "fenny"}" //是一个map
		return c.JSON(c.AllParams())
		// ...
	})

	// GET http://example.com/user/fenny/123
	app.Get("/user/*", func(c *fiber.Ctx) error {
		c.AllParams() // "{"*1": "fenny/123"}"
		return c.JSON(c.AllParams())
		// ...
	})

	app.Listen(":8080")
}
