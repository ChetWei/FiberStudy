package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

//cookie值加密的中间件
func main() {
	app := fiber.New()

	app.Use(encryptcookie.New(encryptcookie.Config{
		// Base64 encoded unique key to encode & decode cookies.
		// Required. Key length should be 32 characters.
		// You may use `encryptcookie.GenerateKey()` to generate a new key.
		Key: "12345678901234567890123456789012",

		/*Custom function to encrypt cookies.
		Optional. Default: EncryptCookie*/

		//Encryptor func(decryptedString, key string) (string, error)

		/*Custom function to decrypt cookies.
		Optional. Default: DecryptCookie*/

		//Decryptor func(encryptedString, key string) (string, error)
	}))

	// Get / 获得解密后的cookie值
	app.Get("/getcookie", func(c *fiber.Ctx) error {

		return c.SendString("value=" + c.Cookies("token"))
	})

	//// Post / 创建加密cookie
	app.Get("/enc", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:  "token",
			Value: "123453424325235235",
		})
		return nil
	})

	app.Listen(":8080")
}
