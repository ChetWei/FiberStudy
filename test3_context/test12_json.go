package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func test() {
	fmt.Println("test() func")
}

func main() {

	app := fiber.New()

	//定义一个结构体类
	type SomeStruct struct {
		Name string
		Age  uint8
	}

	//返回json结构的数据
	app.Get("/json", func(c *fiber.Ctx) error {
		// Create data struct:
		data := SomeStruct{
			Name: "Grame",
			Age:  20,
		}

		//1.将结构体对象编码json结构的数据
		return c.JSON(data)
		// => Content-Type: application/json
		// => "{"Name": "Grame", "Age": 20}"

		//2.将map对象解析成json的数据
		//return c.JSON(fiber.Map{
		//	"name": "Grame",
		//	"age": 20,
		//})
		// => Content-Type: application/json
		// => "{"name": "Grame", "age": 20}"
	})

	//使用jsonp，嵌入回调方法名称，交给客户端
	app.Get("/jsonp", func(c *fiber.Ctx) error {
		// Create data struct:
		data := SomeStruct{
			Name: "Grame",
			Age:  20,
		}

		return c.JSONP(data, "test")
		// 客户端收到的内容=> test({"name": "Grame", "age": 20})

	})

	app.Listen(":8080")
}
