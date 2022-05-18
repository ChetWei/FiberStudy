package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"time"
)

//自定义的时间类型
type CustomTime time.Time

// String() 返回string类型的时间字符串
func (ct *CustomTime) String() string {
	t := time.Time(*ct).String()
	return t
}

//时间格式化
var timeConverter = func(value string) reflect.Value {
	fmt.Println("timeConverter", value)
	if v, err := time.Parse("2006-01-02", value); err == nil {
		return reflect.ValueOf(v)
	}
	return reflect.Value{}
}

func main() {
	app := fiber.New()

	customTime := fiber.ParserType{
		Customtype: CustomTime{}, //使用自定义的类型
		Converter:  timeConverter,
	}

	fiber.SetParserDecoder(fiber.ParserConfig{
		IgnoreUnknownKeys: true,
		ParserType:        []fiber.ParserType{customTime},
		ZeroEmpty:         true,
	})

	type Demo struct {
		Date  CustomTime `form:"date" query:"date"`
		Title string     `form:"title" query:"title"`
		Body  string     `form:"body" query:"body"`
	}

	app.Post("/body", func(c *fiber.Ctx) error {
		var d Demo
		c.BodyParser(&d)
		fmt.Println("d.Date", d.Date.String())
		return c.JSON(d)
	})

	app.Get("/query", func(c *fiber.Ctx) error {
		var d Demo
		c.QueryParser(&d)
		fmt.Println("d.Date", d.Date.String())
		return c.Status(fiber.StatusOK).JSON(d)
	})

	// curl -X POST -F title=title -F body=body -F date=2021-10-20 http://localhost:3000/body

	// curl -X GET "http://localhost:3000/query?title=title&body=body&date=2021-10-20"

	app.Listen(":8080")
}
