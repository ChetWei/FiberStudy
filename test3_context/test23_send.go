package main

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"net/url"
)

//返回消息，消息内容，状态码

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		//return c.Send([]byte("Hello, World!")) // => "Hello, World!"

		//return c.SendString("Hello, World!")

		return c.SendStream(bytes.NewReader([]byte("Hello, World!")))
	})

	app.Get("/file", func(c *fiber.Ctx) error {
		//return c.SendFile("../static/index.html");

		// Disable compression
		return c.SendFile("../static/index.html", false)
	})

	//避免转义 解析文件名称
	app.Get("/file-with-url-chars", func(c *fiber.Ctx) error {
		return c.SendFile(url.PathEscape("hash_sign_#.txt"))
	})

	//
	app.Get("/not-found", func(c *fiber.Ctx) error {

		//设置响应头
		c.Set("Content-Type", "text/plain")

		//发送状态码
		//return c.SendStatus(415)
		// => 415 "Unsupported Media Type"

		c.SendString("Hello, World!") //先发送内容
		return c.SendStatus(415)      //设置状态码
		//=> 415 "Hello, World!"
	})

	app.Listen(":8080")
}
