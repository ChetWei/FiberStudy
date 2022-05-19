package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
)

func main() {

	app := fiber.New()

	// Default middleware config
	//app.Use(requestid.New())

	// Or extend your config for customization
	app.Use(requestid.New(requestid.Config{
		Next: nil,
		//添加响应头 内容即为生成的requestid
		Header: fiber.HeaderXRequestID, //这里使用fiber的常量
		//生成唯一的标识
		Generator: func() string {
			return utils.UUID()
		},
		//定义在本地存储request ID的key
		ContextKey: "requestid",
	}))

	app.Listen(":8080")
}
