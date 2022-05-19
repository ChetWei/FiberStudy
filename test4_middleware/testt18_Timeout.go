package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"time"
)

//超时中间件包装了一个 fiber.Handler方法处理 超时
//如果方法的在给定的处理时间内没有返回，那么超时错误就会发送到ErrorHandler 中心
//the timeout error is set and forwarded to the centralized ErrorHandler

func main() {

	app := fiber.New()

	handler := func(ctx *fiber.Ctx) error {
		time.Sleep(10 * time.Second)
		err := ctx.SendString("Hello world!")
		if err != nil {
			return err
		}
		return nil
	}

	//就会跳转到 hander方法进行处理,允许的处理时长是5秒钟，超过了就进入ErrorHandler
	app.Get("/foo", timeout.New(handler, 5*time.Second))

	app.Listen(":8080")
}
