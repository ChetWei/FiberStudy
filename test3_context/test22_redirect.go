package main

import "github.com/gofiber/fiber/v2"

//重定向

func main() {

	app := fiber.New()

	app.Get("/coffee", func(c *fiber.Ctx) error {
		//302重定向,告知客户端请求下一个地址
		return c.Redirect("/teapot", 302)
	})

	app.Get("/teapot", func(c *fiber.Ctx) error {
		//设置返回的状态码,返回原始状态码
		return c.Status(fiber.StatusOK).Send([]byte("🍵 short and stout 🍵"))
	})

	/*************************************************************************/
	//
	app.Get("/", func(c *fiber.Ctx) error {
		// /user/fiber
		return c.RedirectToRoute("user", fiber.Map{
			"name": "fiber",
		})
	})

	app.Get("/with-queries", func(c *fiber.Ctx) error {
		// /user/fiber?data[0][name]=john&data[0][age]=10&test=doe
		return c.RedirectToRoute("user", fiber.Map{
			"name": "fiber",
			//查询参数
			"queries": map[string]string{"data[0][name]": "john", "data[0][age]": "10", "test": "doe"},
		})
	})

	app.Get("/user/:name", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("name"))
	}).Name("user")

	/*************************************************************************/
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Home page")
	})

	app.Get("/back", func(c *fiber.Ctx) error {
		//重定向
		return c.RedirectBack("/")
	})

	app.Listen(":8080")
}
