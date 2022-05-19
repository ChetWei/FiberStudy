package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

//基础认证

func main() {

	app := fiber.New()

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "john" && pass == "doe" {
				return true
			}
			if user == "admin" && pass == "123456" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.SendFile("../static/unauthorized.html")
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("ok!")
	})

	app.Listen(":8080")
}
