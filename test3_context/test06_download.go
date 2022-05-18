package main

import "github.com/gofiber/fiber/v2"

//下载文件

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		//return c.Download("../static/photo.jpeg"); //不进行重命名，使用源文件名称
		return c.Download("../static/photo.jpeg", "myphoto.jpeg")
	})

	app.Listen(":8080")
}
