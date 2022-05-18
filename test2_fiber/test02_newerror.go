package main

import "github.com/gofiber/fiber/v2"

func main() {
	//返回错误消息
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {

		return fiber.NewError(789, "自定义错误消息")
	})
	//启动应用
	app.Listen("127.0.0.1:8080")
}
