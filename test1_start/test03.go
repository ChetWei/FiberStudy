package main

import "github.com/gofiber/fiber/v2"

func main() {

	//一个fiber的应用程序
	app := fiber.New()

	//请求根路径 http://127.0.0.1:8080/hello
	app.Get("/:name?", func(c *fiber.Ctx) error {
		//把获取到的参数绑定到name里面
		if c.Params("name") != "" {
			return c.SendString("value:" + c.Params("name"))
		}
		return c.SendString("参数为空！")
	})
	//启动应用
	app.Listen("127.0.0.1:8080")

}
