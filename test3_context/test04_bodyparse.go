package main

import (
	"github.com/gofiber/fiber/v2"
)

//处理客户端的请求体

func main() {

	app := fiber.New()

	// 返回原始请求数据
	// curl -X POST http://localhost:8080 -d user=john
	app.Post("/", func(c *fiber.Ctx) error {

		return c.Send(c.Body()) // []byte("user=john")
	})

	//根据请求体的数据类型进行解析
	/*
		application/x-www-form-urlencoded 	form
		multipart/form-data 				form
		application/json					json
		application/xml						xml
		text/xml							xml
	*/

	// 字段名称必须大写
	type Person struct {
		//可以接受三种类型的数据
		Name string `json:"name" xml:"name" form:"name"`
		Pass string `json:"pass" xml:"pass" form:"pass"`
	}

	app.Post("/data", func(c *fiber.Ctx) error {
		p := new(Person) //创建一个结构体对象
		//将解析的内容赋值给结构体对象
		if err := c.BodyParser(p); err != nil {
			return err
		}
		return c.JSON(p)
	})
	// curl -X POST -H "Content-Type: application/json" --data "{\"name\":\"john\",\"pass\":\"doe\"}" localhost:3000
	// curl -X POST -H "Content-Type: application/xml" --data "<login><name>john</name><pass>doe</pass></login>" localhost:3000
	// curl -X POST -H "Content-Type: application/x-www-form-urlencoded" --data "name=john&pass=doe" localhost:3000
	// curl -X POST -F name=john -F pass=doe http://localhost:3000
	// curl -X POST "http://localhost:3000/?name=john&pass=doe"

	app.Listen(":8080")

}
