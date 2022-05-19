package main

import (
	"expvar"
	"fmt"
	"github.com/gofiber/fiber/v2"
	expvarmw "github.com/gofiber/fiber/v2/middleware/expvar"
)

//使用expvar定义一个变量
var count = expvar.NewInt("count")

func main() {

	app := fiber.New()

	app.Use(expvarmw.New(
		expvarmw.Config{
			Next: nil,
		},
	))

	app.Get("/", func(c *fiber.Ctx) error {
		count.Add(1)
		return c.SendString(fmt.Sprintf("hello expvar count %d", count.Value()))
	})

	//127.0.0.1:8080/debug/vars 来访问所有暴露的值

	app.Listen(":8080")
}
