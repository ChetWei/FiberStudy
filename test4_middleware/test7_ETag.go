package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

//ETag中间件，web服务不需要重复发送没有改变的响应数据
//让缓存更加高效，可以节省带宽

func main() {

	app := fiber.New()

	app.Use(etag.New(etag.Config{
		Next: nil,
		// Weak indicates that a weak validator is used. Weak etags are easy
		// to generate, but are far less useful for comparisons. Strong
		// validators are ideal for comparisons but can be very difficult
		// to generate efficiently. Weak ETag values of two representations
		// of the same resources might be semantically equivalent, but not
		// byte-for-byte identical. This means weak etags prevent caching
		// when byte range requests are used, but strong etags mean range
		// requests can still be cached.
		Weak: false,
	}))

	// 会在请求头中添加  Etag: "W/"13-1831710635"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8080")
}
