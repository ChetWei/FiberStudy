package main

import "github.com/gofiber/fiber/v2"

/*
accept
*/

func main() {
	app := fiber.New()

	// Accept-Charset: utf-8, iso-8859-1;q=0.2
	// Accept-Encoding: gzip, compress;q=0.2
	// Accept-Language: en;q=0.8, nl, ru

	app.Get("/", func(c *fiber.Ctx) error {
		c.AcceptsCharsets("utf-16", "iso-8859-1")
		// "iso-8859-1"

		c.AcceptsEncodings("compress", "br")
		// "compress"

		c.AcceptsLanguages("pt", "nl", "ru")
		// "nl"
		// ...
		return c.SendString("OK!")
	})

	app.Listen(":8080")
}
