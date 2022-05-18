package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//获取请求头的相关信息

func main() {
	app := fiber.New()

	app.Get("/req", func(c *fiber.Ctx) error {
		//TODO 为什么这里获取的是空?
		var contentType1 string = c.Get("Content-Type", "application/json") // ""
		//获取请求头信息
		var contentType2 string = c.GetRespHeader("Content-Type") // "text/plain"
		fmt.Println("contentType1: ", contentType1)
		fmt.Println("contentType2: ", contentType2)
		return nil
	})

	//获取响应头信息
	app.Get("/resp", func(c *fiber.Ctx) error {
		//TODO 获取不到?
		r1 := c.GetRespHeader("X-Request-Id")      // "8d7ad5e3-aaf3-450b-a241-2beb887efd54"
		r2 := c.GetRespHeader("Content-Type")      // "text/plain"
		r3 := c.GetRespHeader("something", "john") // "john"
		// ..
		fmt.Println(r1)
		fmt.Println(r2)
		fmt.Println(r3)
		return nil
	})

	app.Listen(":8080")
}
