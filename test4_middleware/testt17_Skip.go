package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"
)

//跳过指定的方法，如果设置值为true
//Skip middleware for Fiber that skips a wrapped handler， if a predicate is true.

func main() {

	app := fiber.New()

	handler := func(ctx *fiber.Ctx) error {
		err := ctx.SendString("Hello, World 👋!")
		if err != nil {
			return err
		}
		return nil
	}

	app.Use(skip.New(handler, func(ctx *fiber.Ctx) bool {
		return ctx.Method() == fiber.MethodOptions
	}))

	app.Listen(":8080")
}
