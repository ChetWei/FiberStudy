package main

import "github.com/gofiber/fiber/v2"

//在响应头中设置Link字段

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		//在响应头中设置Link字段
		c.Links(
			"http://api.example.com/users?page=2", "next",
			"http://api.example.com/users?page=5", "last",
		)
		// Link: <http://api.example.com/users?page=2>; rel="next",
		//       <http://api.example.com/users?page=5>; rel="last"

		// ...
		return nil
	})

	app.Listen(":8080")
}
