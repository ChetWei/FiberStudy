package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//跨域配置

func main() {
	app := fiber.New()

	// Default config
	//app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		//next跳过这个跨域中间件
		Next: nil,
		//允许跨域访问的域名   * 表示允许所有
		AllowOrigins: "https://gofiber.io, https://gofiber.net",
		//允许跨域访问的方法
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		// AllowCredentials indicates whether or not the response to the request
		// can be exposed when the credentials flag is true. When used as part of
		// a response to a preflight request, this indicates whether or not the
		// actual request can be made using credentials.
		AllowCredentials: false,
		// ExposeHeaders defines a whitelist headers that clients are allowed to access
		ExposeHeaders: "",
		// MaxAge indicates how long (in seconds) the results of a preflight request  can be cached.
		MaxAge: 0,
	}))

	app.Listen(":8080")
}
