package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New()

	//ListenTLS 使用 certfile 和 keyfile 路径从给定地址提供 HTTPS 请求，以作为 TLS 证书和密钥文件。
	app.ListenTLS(":443", "./cert.pem", "./cert.key")
}
