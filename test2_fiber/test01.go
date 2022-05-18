package main

import "github.com/gofiber/fiber/v2"

func main() {
	//自定义配置
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
	})

	// ...
	//启动应用
	app.Listen("127.0.0.1:8080")
}
