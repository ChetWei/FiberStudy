package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)
import "github.com/gofiber/fiber/v2/middleware/recover"

//Recover middleware for Fiber that recovers from panics anywhere in the stack chain
//and handles the control to the centralized ErrorHandler.

/*
恢复 Fiber 的中间件，可从堆栈链中的任何地方的panic中恢复，并将控制权交给集中式 ErrorHandler。
*/
func defaultStackTraceHandler(c *fiber.Ctx, e interface{}) {
	fmt.Println("do defaultStackTraceHandler")
}

func main() {

	app := fiber.New()

	app.Use(recover.New(recover.Config{
		Next: nil,
		//启用处理堆栈跟踪
		EnableStackTrace: true,
		//处理的方法
		StackTraceHandler: defaultStackTraceHandler,
	}))

	//panic 会被 recover中间件捕获
	app.Get("/", func(c *fiber.Ctx) error {
		panic("I'm an error")
	})

	app.Listen(":8080")
}
