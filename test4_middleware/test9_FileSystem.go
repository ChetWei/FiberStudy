package main

import (
	"github.com/gobuffalo/packr/v2" //是一个将静态文件嵌入到 Go 二进制文件中的工具
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"log"
)

//Filesystem middleware
// 可以 让一个文件夹提供文件服务

func main() {

	app := fiber.New()

	//app.Use(filesystem.New(filesystem.Config{
	//	Next: nil,
	//	Root:         http.Dir("../static"), //将一个目录作为根路径
	//	Browse:       true, //允许浏览
	//	Index:        "index.html",
	//	NotFoundFile: "404.html",
	// The value for the Cache-Control HTTP-header
	// that is set on the file response. MaxAge is defined in seconds.
	//	MaxAge:       3600,
	//}))

	app.Use("/static", filesystem.New(filesystem.Config{
		Root: packr.New("Assets Box", "../static"),
	}))

	log.Fatal(app.Listen(":8080"))

}
