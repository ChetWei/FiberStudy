# FiberStudy
gofiber framework study



## 1.开始

初始化项目模块，下载安装gofiber

```go
go mod init <moudlue_name>

go get github.com/gofiber/fiber/v2
```

### 启动一个简单的demo

```go
import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	//一个fiber的实例
	app := fiber.New()

	//请求根路径，GET方法，"/"虚拟路径，
  //func(*fiber.Ctx) error 是包含当路由匹配时执行的 Context 的回调函数。
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//启动应用
	app.Listen("127.0.0.1:8080")
}
```

### 基本路由

```go
// Respond with "Hello, World!" on root path, "/"
app.Get("/", func(c *fiber.Ctx) error {
  return c.SendString("Hello, World!") //返回string字符串
})
```

### 参数路由

```go
// GET http://localhost:8080/hello world

app.Get("/:value", func(c *fiber.Ctx) error {
  return c.SendString("value: " + c.Params("value"))
  // => Get request with value: hello world
})
```

`/:value`可以获取`/`路径后面的值,但是如果是多级路径，就无法获取

### 可选参数

```go
app.Get("/:name?", func(c *fiber.Ctx) error {
		//把获取到的参数绑定到name里面
		if c.Params("name") != "" {
			return c.SendString("value:" + c.Params("name"))
		}
		return c.SendString("参数为空！")
	})
```

`/:name?`具有上面的功能，同时可以匹配 `/`路径，也就是/路径后面没有参数的情况

### 通配符

```go
//获取 /api后面的所有值
	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
	})
```

### 访问静态文件

```go
func main() {
	app := fiber.New()

  //在static目录下有一张 photo.jpeg图片
	app.Static("/", "../static")
	app.Listen(":8080")

}
```

请求 http://127.0.0.1:8080/photo.jpeg



## 2.Fiber的配置

```go
//自定义配置
app := fiber.New(fiber.Config{
    Prefork:       true,
    CaseSensitive: true,
    StrictRouting: true,
    ServerHeader:  "Fiber",
    AppName: "Test App v1.0.1"
})

// ...
```

### 配置字段

- Prefork (预处理)

  > 默认false，使用 [`SO_REUSEPORT`](https://lwn.net/Articles/542629/)套接字选项。这将产生多个 Go 进程，监听同一端口。**如果启用，应用程序将需要通过 shell 运行，因为预工作模式设置环境变量。**

- ServerHeader (服务器头)

  > 默认为空，使用给定值启用`服务器` HTTP 头。

- StrictRouting (严格路由)

  > 默认false，启用时，路由器将 `/foo` 和 `/foo/` 视为不同。否则，路由器将 `/Foo` 和 `/foo/` 视为相同

- CaseSensitive (区分大小写)

  > 默认false，启用时，`/Foo` 和 `/Foo` 是不同的路由。禁用时，`/Foo` 和 `/foo` 视为相同

- Immutable (不可变)

  > 默认false，启用时，上下文方法返回的所有值都是不可变的。

- UnescapePath (未解除路径）

  > 默认false，在设置上下文路径之前将路由中的所有编码字符转换回，以便路由也可以使用 URL 编码的特殊字符

- ETag（E 标签）

  > 默认false,启用或禁用 ETag 头生成，因为弱 ETag 和强 ETag 都是使用相同的哈希方法 (CRC-32) 生成的。弱 ETag 是启用时的默认值

- BodyLimit（Body 限制）

  > 默认`4x1024x1024`，设置请求正文允许的最大大小，如果大小超过配置的限制，则发送 `413 - Request Entity Too Large` 响应

- Concurrency（并发）

  > 默认256x1024 ,最大并发连接数。 int 

- Views（视图）

  > 默认 nil，视图是包装渲染功能的界面。有关支持的引擎,参考模板中间件 Views

- ReadTimeout（读取时间）

  > 允许读取完整请求 (包括正文) 的时间量。默认超时是无限制的 time.Duration

- WriteTimeout（写入时间）

  > 响应超时写入之前的最大持续时间。默认超时是无限制的 time.Duration

- IdleTimeout（闲置时间）

  > 启用保持活动状态时等待下一个请求的最长时间。如果 IdleTimeout 为零，则使用 ReadTimeout 的值。默认 nil

- ReadBufferSize（读取缓冲大小）

  > 每个请求读取的连接缓冲区大小。这也限制了最大标题大小。如果您的客户端发送多 KB 请求 URI 和 / 或多 KB 标头 (例如，大 cookie)，请增加此缓冲区。默认4096

- WriteBufferSize （写入缓冲大小）

  > 响应写入的每个连接缓冲区大小。 默认4096

- CompressedFileSuffix （压缩文件后缀）

  > 为原始文件名添加后缀，并尝试以新文件名保存生成的压缩文件。默认string`.fiber.gz`

- ProxyHeader（代理头）

  > 这将使 `c.IP()` 能够返回给定标头键的值。默认情况下，`c.IP()` 将从 TCP 连接返回远程 IP，如果您在负载平衡器后面，例如 **X-Forwarded-**，则此属性可能很有用。默认string `""`

- GETOnly （仅限 GET）

  > 如果设置为 true，则拒绝所有非 GET 请求。对于仅接受 GET 请求的服务器，此选项作为防 DoS 保护非常有用。如果设置了 GETOnly，则请求大小受 ReadBufferSize 限制。默认 `false`

- ErrorHandler（错误请求头）

  > 当 fiber 返回错误时，执行 ErrorHandler。处理程序。安装的 fiber 错误处理程序由顶级应用程序保留，并应用于前缀相关请求。 默认 `DefaultErrorHandler`

- DisableKeepalive（关闭连接）

  > 禁用保持活动连接，服务器将在向客户端发送第一个响应后关闭传入连接。默认`false`

- DisableDefaultDate （关闭默认时间）

  > 当设置为 true 时，将导致从响应中排除默认日期头。默认`false`

- DisableDefaultContentType（关闭默认内容类型）

  >设置为 true 时，将导致从响应中排除默认内容类型标题。默认`false`

- DisableHeaderNormalizing（关闭头正常化）

  > 默认情况下，所有标题名称都是标准化的：内容类型 -> 内容类型。默认`false`

- DisableStartupMessage（关闭启动信息)

  > 设置为 true 时，它不会打印调试信息 。默认`false`

- AppName (应用名)

  > AppName (应用名) 默认 string `""`

- EnableTrustedProxyCheck（启用可信任代理检查）

  >设置为 true 时，fiber 将使用 TrustedProxy 列表检查代理是否受信任。
  >
  >默认情况下，c.Protocol() 将从 X-Forwarded-Proto、X-Forwarded-Protocol、X-Forwarded-Ssl 或 X-Url-Scheme 头获取值，c.IP() 将从代理头获取值，c.Hostname() 将从 X-Forwarded-Host 头获取值。如果 EnableTrustedProxyCheck 为 true，并且 RemoteIP 在受信任代理的列表中 c.Protocol()，c.IP() 和 c.Hostname() 将在禁用受信任代理时具有相同的行为，如果 RemoteIP 不在列表中，c.Protocol() 将在应用程序处理 tls 连接时返回 https，否则，c.IP() 将从 fasthttp 上下文返回 RemoteIP () ，c.Hostname() 将返回 fasthttp.Request.URI().Host()。默认`false`

- TrustedProxies（可信任代理）

  > 包含受信任代理 IP 的列表。请查看`启用受信任代理检查`文档。它可以采用 IP 或 IP 范围地址。如果它得到 IP 范围，它将迭代所有可能的地址

- DisablePrepareSMultipartForm （关闭复合文件表单)

  > 如果设置为 true，则不会预解析多部分表单数据。此选项对于希望将多部分表单数据视为二进制 blob 或选择何时解析数据的服务器非常有用。默认`false`

- StreamRequestBody（流媒体请求内容）

  > StreamRequestBody 启用请求正文流，并在给定正文大于当前限制时更快调用处理程序。默认`false`



### 返回错误消息

```go
app.Get("/", func(c *fiber.Ctx) error {
    return fiber.NewError(782, "Custom error message")
})
```

### 判断当前进程是否是预处理的结果

```go
//预处理将产生子进程
app := fiber.New(fiber.Config{
    Prefork: true,
})

if !fiber.IsChild() {
    fmt.Println("I'm the parent process")
} else {
    fmt.Println("I'm a child process")
}

// ...
```

## App实例

### Static 配置

> 使用 **Static** 方法来服务静态文件，例如 **images**，**CSS,** 和 **JavaScript**。

```go
app.Static("/", "./public")
```

#### 静态文件的控制

