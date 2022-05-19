package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"log"
	"time"
)

//Session的处理，支持拓展数据库，默认是在内存中
/*
func New(config ...Config) *Store
func (s *Store) RegisterType(i interface{})
func (s *Store) Get(c *fiber.Ctx) (*Session, error)
func (s *Store) Reset() error

func (s *Session) Get(key string) interface{}
func (s *Session) Set(key string, val interface{})
func (s *Session) Delete(key string)
func (s *Session) Destroy() error
func (s *Session) Regenerate() error
func (s *Session) Save() error
func (s *Session) Fresh() bool
func (s *Session) ID() string
func (s *Session) Keys() []string
*/

func main() {

	app := fiber.New()

	// 存储所有应用的会话
	//设置之后每次请求都会在cookie中添加 session_id 和 token
	store := session.New(session.Config{
		// session的过期时间
		// Optional. Default value 24 * time.Hour
		Expiration: time.Minute * 1,

		//session的key
		// Optional. Default value "session_id".
		CookieName: "session_id",

		// 生成cookie的方式
		// Optional. Default value utils.UUID
		KeyGenerator: utils.UUID,

		// 存储session的接口
		// Optional. Default value memory.New()
		//Storage: fiber.Storage,

		// Domain of the cookie.
		// Optional. Default value "".
		//CookieDomain string

		// Path of the cookie.
		// Optional. Default value "".
		//CookiePath string

		// Indicates if cookie is secure.
		// Optional. Default value false.
		//CookieSecure bool

		// Indicates if cookie is HTTP only.
		// Optional. Default value false.
		//CookieHTTPOnly bool

		// Sets the cookie SameSite attribute.
		// Optional. Default value "Lax".
		//CookieSameSite string

	})

	app.Get("/login", func(ctx *fiber.Ctx) error {
		// 从内存中获取session
		sess, err := store.Get(ctx)
		if err != nil {
			panic(err)
		}

		name := ctx.Query("name", "nobody")

		// 设置 key/value
		sess.Set("name", name)

		//保存session的值,以保证其他的方法可以拿到
		if err := sess.Save(); err != nil {
			panic(err)
		}

		return ctx.SendString(fmt.Sprintf("login success! %v", name))
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		// 从内存中获取session
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		name := sess.Get("name")
		fmt.Println(name)
		//删除key
		sess.Delete("name")

		// 销毁当前会话
		if err := sess.Destroy(); err != nil {
			panic(err)
		}

		return c.SendString(fmt.Sprintf("Bye %v", name))
	})

	app.Get("/home", func(c *fiber.Ctx) error {
		// 从内存中获取session
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		// 获取 value
		nameV := sess.Get("name")
		if nameV == nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// 获取所有的key
		keys := sess.Keys()
		log.Println(keys)

		return c.SendString(fmt.Sprintf("Welcome %v", nameV))
	})

	app.Listen(":8080")
}
