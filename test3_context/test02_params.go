package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//params参数获取  /路径后面的参数

func main() {
	app := fiber.New()

	// GET http://example.com/user/fenny
	app.Get("/user/:name", func(c *fiber.Ctx) error {

		c.AllParams() // "{"name": "fenny"}" //是一个map

		name := c.Params("name")

		fmt.Println(name)
		return c.JSON(c.AllParams())
		// ...
	})

	// GET http://example.com/user/fenny/123
	app.Get("/user/*", func(c *fiber.Ctx) error {

		c.AllParams() // "{"*1": "fenny/123"}"

		c.Params("*")  // "fenny/123"
		c.Params("*1") // "fenny/123"

		return c.JSON(c.AllParams())
		// ...
	})

	//多个通配符的使用
	app.Get("/v1/*/shop/*", func(c *fiber.Ctx) error {
		// GET:   /v1/brand/4/shop/blue/xs

		c.Params("*1") // "brand/4"
		c.Params("*2") // "blue/xs"
		return nil

	})

	//解析数据类型
	// GET http://example.com/user/123
	app.Get("/user/:id", func(c *fiber.Ctx) error {
		//将获取的参数解析为int类型
		id, err := c.ParamsInt("id") // int 123 and no error
		if err == nil {

		}
		fmt.Println(id)
		return nil
	})
	app.Listen(":8080")
}
