package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

//获取通过 地址栏 ? 请求的参数

func main() {

	app := fiber.New()

	// GET http://example.com/?name=desc&password=nike
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c.Query("name"))             //desc
		fmt.Println(c.Query("password", "nike")) // "nike"

		return nil
	})

	//解析query数据到结构体中
	// 字段名称必须大写，通过标签绑定请求字段
	type Person struct {
		Name     string   `query:"name"`
		Pass     string   `query:"pass"`
		Products []string `query:"products"` // &products=shoe,hat 两个值 数组
	}

	app.Get("/parse", func(c *fiber.Ctx) error {

		p := new(Person)

		if err := c.QueryParser(p); err != nil {
			return err
		}

		log.Println(p.Name)     // john
		log.Println(p.Pass)     // doe
		log.Println(p.Products) // [shoe, hat]
		return nil
	})
	// Run tests with the following curl command
	//
	// curl "http://localhost:3000/?name=john&pass=doe&products=shoe,hat"

	app.Listen(":8080")
}
