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
		file, _ := c.FormFile("document") //获取第一个文件
		formValue := c.FormValue("desc")  //获取一个非文件字段的内容
		fmt.Println(formValue)
		// 保存
		return c.SaveFile(file, fmt.Sprintf("../upload/%s", file.Filename))
	})

	//保存多个上传的文件
	app.Post("/many", func(c *fiber.Ctx) error {
		// Parse the multipart form:
		if form, err := c.MultipartForm(); err == nil {
			// => *multipart.Form
			if token := form.Value["token"]; len(token) > 0 {
				// Get key value:
				fmt.Println(token[0])
			}

			// 获取所有key 为documents的文件:
			files := form.File["documents"]
			// => []*multipart.FileHeader

			// 遍历这些文件:
			for _, file := range files {
				//打印文件的名称，大小(B)，类型
				fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
				// => "tutorial.pdf" 360641 "application/pdf"

				// 保存文件:
				if err := c.SaveFile(file, fmt.Sprintf("../upload/%s", file.Filename)); err != nil {
					return err
				}
			}
		}

		return c.SendString("ok")
	})

	app.Listen(":8080")

}
