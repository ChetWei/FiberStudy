package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

//将请求代理到多个服务器

func main() {

	app := fiber.New()

	// 由服务器请求内容，返回给客户端
	app.Get("/baidu", proxy.Forward("https://www.baidu.com"))

	// http://localhost:8080/cAwa71N

	app.Get("/:id", func(c *fiber.Ctx) error {
		url := "https://i.imgur.com/gallery/" + c.Params("id") + ".gif"
		if err := proxy.Do(c, url); err != nil {
			return err
		}
		// Remove Server header from response
		c.Response().Header.Del(fiber.HeaderServer)
		return nil
	})

	// app.use 之前的内容不受影响

	// 自定义负载均衡，将请求代理到其他服务器
	app.Use(proxy.Balancer(proxy.Config{
		Next: nil,
		//使用的是轮训方式
		// Servers defines a list of <scheme>://<host> HTTP servers,
		Servers: []string{
			"http://localhost:8081",
			"http://localhost:8082",
		},
		//转发之前修改 请求
		ModifyRequest: func(c *fiber.Ctx) error {
			c.Request().Header.Add("X-Real-IP", c.IP())
			return nil
		},
		//修改响应内容
		ModifyResponse: func(c *fiber.Ctx) error {
			c.Response().Header.Del(fiber.HeaderServer)
			return nil
		},
	}))

	app.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world！")
	})

	app.Listen(":8080")
}
