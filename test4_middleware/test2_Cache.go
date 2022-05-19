package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/utils"
	"time"
)

//拦截响应，并且缓存数据。缓冲body，content-type,以及statuscode，根据C.Path() 作为唯一标志

func main() {

	app := fiber.New()

	// 配置cache
	app.Use(cache.New(cache.Config{
		//如果为true，下一个方法跳过 这个缓存中间件
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		//过期时间
		Expiration: 1 * time.Minute,
		// 允许客户端缓存
		CacheControl: true,
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(c.Path())
		},
		ExpirationGenerator:  nil,
		StoreResponseHeaders: false,
		Storage:              nil,
	}))

	app.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world~")
	})

	app.Listen(":8080")
}
