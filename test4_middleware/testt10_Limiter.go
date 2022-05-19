package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

//限制客户端对接口的重复请求,根据ip来统计的

func main() {

	app := fiber.New()

	app.Use(limiter.New(limiter.Config{
		//为true跳过这个中间件
		Next: func(c *fiber.Ctx) bool {
			//本地请求不限制
			//return c.IP() == "127.0.0.1"
			return false
		},
		//在Expiration阶段内的最大连接数量
		Max: 5,
		//在内存中记录链接请求 时间段内
		Expiration: 30 * time.Second,
		//KeyGenerator allows you to generate custom keys, by default c.IP() is used
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		// 请求触发限制条件的时候调用
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendFile("../static/toofast.html")
		},
		//如果设置为true，状态码 StatusCode >= 400 的请求 不会被统计数量
		SkipFailedRequests: false,
		//如果设置为true，状态码 StatusCode < 400 的请求 不会被统计数量
		SkipSuccessfulRequests: false,

		//LimiterMiddleware is the struct that implements limiter middleware.
		//LimiterMiddleware: FixedWindow{},

		// Store is used to store the state of the middleware
		// Default: an in memory store for this process only
		//Storage: nil,
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world!")
	})

	app.Listen(":8080")
}
