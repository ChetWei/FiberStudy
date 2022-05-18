package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Static("/", "../static")
	app.Listen(":8080")

	//请求 http://127.0.0.1:8080/photo.jpeg
}
