package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//判断请求客户端的ip等信息

func main() {
	app := fiber.New()

	// GET http://google.com/search

	app.Get("/", func(c *fiber.Ctx) error {
		r1 := c.Hostname() // "google.com"  获取用户请求的域名,没有域名显示ip:port

		// Host: "tobi.ferrets.example.com" //域名分割
		c.Subdomains()  // ["ferrets", "tobi"]
		c.Subdomains(1) // ["tobi"]

		r2 := c.IP() //客户端请求的ip
		r3 := c.IPs()

		fmt.Println(r1)
		fmt.Println(r2)
		fmt.Println(r3)

		//判断conten type 如果请求头是 text/html; charset=utf-8
		fmt.Println(c.Is("html"))  //true
		fmt.Println(c.Is(".html")) //true
		fmt.Println(c.Is("json"))  //false

		//判断是否来自本地的请求
		fmt.Println(c.IsFromLocal())

		return nil
	})

	app.Listen(":8080")
}
