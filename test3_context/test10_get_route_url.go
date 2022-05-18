package main

import "github.com/gofiber/fiber/v2"

//路由重命名，根据路由名称进行调用

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Home page")
	}).Name("home")

	app.Get("/user/:id", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("id"))
	}).Name("user.show")

	app.Get("/test", func(c *fiber.Ctx) error {
		//location 是调用的路由路径
		location, _ := c.GetRouteURL("home", nil)
		//location, _ := c.GetRouteURL("user.show", fiber.Map{"id": 1})
		return c.SendString(location)
	})

	// /test returns "/user/1"

	app.Listen(":8080")
}
