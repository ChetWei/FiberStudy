package main

import "github.com/gofiber/fiber/v2"

//设置返回的响应头

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.Append("Link", "http://google.com", "http://localhost")
		// => Link: http://localhost, http://google.com

		//相同的字段，会添加进去
		c.Append("Link", "Test")
		// => Link: http://localhost, http://google.com, Test
		//
		return nil
	})

	//设置 下载附件
	app.Get("/a", func(c *fiber.Ctx) error {
		c.Attachment()
		// => Content-Disposition: attachment
		//设置下载附件
		c.Attachment("../static/photo.jpeg")
		// => Content-Disposition: attachment; filename="logo.png"
		// => Content-Type: image/png

		// ...
		return nil
	})

	app.Listen(":8080")
}
