//监控服务器性能指标，测试版本，api可能修改

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New()

	app.Get("/dashboard", monitor.New())

	log.Fatal(app.Listen(":8080"))
}
