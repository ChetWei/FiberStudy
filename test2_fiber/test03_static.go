package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	app := fiber.New()

	app.Static("/", "../static", fiber.Static{
		////设置为true时，服务器尝试通过缓存压缩文件最小化CPU使用率。
		//    //这与Github.com/gofiber/Compression中间件不同。
		//    //可选，默认值False.
		Compress: true,
		//设置为true时，启用字节范围请求。
		//可选，默认值False.
		ByteRange: true,
		//设置为true时，启用目录浏览。
		//可选，默认值false。
		Browse: true,
		//用于服务目录的索引文件的名称。
		//可选，默认值“index.html”。
		Index: "index.html",
		//非活动文件处理程序的失效持续时间。
		//使用否定时间。要禁用它。
		// 可选，默认值10 * time.second。
		CacheDuration: 19 * time.Second,
		//缓存控制http-head的值
		//设置在文件响应上。 maxage在几秒钟内定义。
		MaxAge: 3600,
	})
	app.Listen(":8080")

	//请求 http://127.0.0.1:8080/photo.jpeg
}
