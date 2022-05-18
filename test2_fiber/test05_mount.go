package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

//Mount挂载

func main() {
	micro := fiber.New()
	micro.Get("/doe", func(c *fiber.Ctx) error {
		return c.SendString("/doe方法")
	})

	app := fiber.New()
	//挂载，将 "/john"的请求，挂载到micro实例上 那么 /john/doe就会到达 micro的 /doe中
	app.Mount("/john", micro) // GET /john/doe -> 200 OK

	log.Fatal(app.Listen(":8080"))
}
