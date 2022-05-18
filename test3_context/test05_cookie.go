package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

//清除客户端的cookie

func main() {

	app := fiber.New()
	//设置cookie字段
	app.Get("/set", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    "randomvalue",
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			SameSite: "lax",
		})
		return c.SendString("set ok!")
	})

	app.Get("/delete", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name: "token",
			// Set expiry date to the past
			Expires:  time.Now().Add(-(time.Hour * 2)),
			HTTPOnly: true,
			SameSite: "lax",
		})
		return c.SendString("delete ok!")
	})

	//ClearCookie将指定的cookie字段设置为过期
	app.Get("/clear", func(c *fiber.Ctx) error {
		// Clears all cookies:
		c.ClearCookie()
		// Expire specific cookie by name:
		c.ClearCookie("user")
		// Expire multiple cookies by names:
		c.ClearCookie("token", "session", "track_id", "version")
		return c.SendString("clear ok!")

	})

	// Create cookie 创建一个cookie对象
	app.Get("/", func(c *fiber.Ctx) error {
		cookie := new(fiber.Cookie)
		cookie.Name = "john"
		cookie.Value = "doe" //set john=doe
		cookie.Expires = time.Now().Add(24 * time.Hour)

		// Set cookie
		c.Cookie(cookie)
		// ...
		return c.SendString("ok!")
	})

	//根据cookie的名称获取值
	app.Get("/get", func(c *fiber.Ctx) error {
		// Get cookie by key:
		a1 := c.Cookies("john", "empty") // "john"
		a2 := c.Cookies("tom", "empty")  // "john"

		fmt.Println("a1 ", a1)
		fmt.Println("a2 ", a2)

		return nil
	})

	app.Listen(":8080")
}
