package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
	"time"
)

//避免 跨站请求伪造

func main() {
	app := fiber.New()

	//
	app.Use(csrf.New(csrf.Config{
		Next:           nil,                   //如果为true，跳过这个中间件
		KeyLookup:      "header:X-Csrf-Token", //KeyLookup 是一个格式为“<source>:<key>”的字符串，用于 从请求中提取令牌。
		CookieName:     "csrf_",               // Name of the session cookie.
		CookieSameSite: "Strict",              //Domain of the CSRF cookie.
		Expiration:     1 * time.Hour,         // Expiration is the duration before csrf token will expire
		KeyGenerator:   utils.UUID,            // KeyGenerator creates a new CSRF token
	}))

	app.Listen(":8080")
}
