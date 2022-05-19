package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

//根据 请求的头 Accept-Encoding，压缩响应数据 使用 gzip, deflate and brotli

func main() {
	app := fiber.New()

	// 默认的中间件
	//app.Use(compress.New())

	// 压缩等级
	//	 Optional. Default: LevelDefault
	// LevelDisabled:         -1
	// LevelDefault:          0
	// LevelBestSpeed:        1
	// LevelBestCompression:  2
	//app.Use(compress.New(compress.Config{
	//	Level: compress.LevelBestSpeed, // 1
	//}))

	// 对特定的路由跳过压缩
	app.Use(compress.New(compress.Config{
		Next: func(c *fiber.Ctx) bool {
			//返回true就不进行压缩
			return c.Path() == "/dont_compress"
		},
		Level: compress.LevelBestSpeed, // 1
	}))

	app.Get("/dont_compress", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world!")
	})

	app.Listen(":8080")
}
