package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"time"
)

//日志记录

func main() {

	app := fiber.New()

	//打开日志文件
	logFile, err := os.OpenFile("./123.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()

	app.Use(logger.New(logger.Config{
		//格式化日志的内容
		Format:     "[${time}] ${status} - ${ip}  ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05-0700",
		TimeZone:   "Asia/Shanghai",
		// 时间更新间隔
		// Optional. Default: 500 * time.Millisecond
		TimeInterval: 500 * time.Millisecond,
		//Output:       logFile, //定义了这个值，就会写入到文件中
		Output: os.Stderr,
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world!")
	})

	app.Listen(":8080")
}
