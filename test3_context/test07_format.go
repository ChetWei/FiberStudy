package main

import "github.com/gofiber/fiber/v2"

//根据请求头的accept来自动格式化内容
func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// Accept: text/plain
		//c.Format("Hello, World!")
		// => Hello, World!

		// Accept: text/html
		//c.Format("Hello, World!")
		// => <p>Hello, World!</p>

		// Accept: application/json
		return c.Format("Hello, World!")
		// => "Hello, World!"
		// ..

	})

	app.Listen(":8080")
}
