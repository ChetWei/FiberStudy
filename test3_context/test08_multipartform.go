package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//获取上传的文件

func main() {

	app := fiber.New()

	//保存用户上传的文件
	app.Post("/", func(c *fiber.Ctx) error {
		// Get first file from form field "document":
		file, _ := c.FormFile("document")
		formValue := c.FormValue("desc") //获取一个非文件字段的内容
		fmt.Println(formValue)
		// 保存
		return c.SaveFile(file, fmt.Sprintf("../upload/%s", file.Filename))
	})

	app.Listen(":8080")

}
