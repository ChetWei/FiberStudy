package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	//一个fiber的应用程序
	app := fiber.New()

	//请求根路径 http://127.0.0.1:8080/hello
	app.Get("/:value", func(c *fiber.Ctx) error {
		//c.Params("value") 得到的值是 hello
		return c.SendString("value:" + c.Params("value"))
	})

	//启动应用
	app.Listen("127.0.0.1:8080")
}
