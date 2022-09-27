package server

import (
	"app/controllers"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// サーバ起動
func StartServer() {
	r := Router()
	r.Run()
}

// ログフォーマッタ
var customLogFormatter = func(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}

	// ここでフォーマットを決めている
	return fmt.Sprintf("[GIN] %v| %s |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		param.Request.Header.Get("X-Request-Id"), // ここに追加
		statusColor,
		param.StatusCode,
		resetColor,
		param.Latency,
		param.ClientIP,
		methodColor,
		param.Method,
		resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

// カスタムログ
func customLogger() gin.LoggerConfig {
	conf := gin.LoggerConfig{
		SkipPaths: []string{"/api/v1/healthcheck"},
	}
	conf.Formatter = customLogFormatter
	return conf
}

// カスタムハンドラ
func customRequestidHandler() requestid.Handler {
	return func(c *gin.Context, requestID string) {
		c.Request.Header.Set("X-Request-Id", requestID)
	}
}

func Router() *gin.Engine {
	r := gin.New()
	r.Use(requestid.New(
		requestid.WithHandler(customRequestidHandler())),
	)
	r.Use(gin.LoggerWithConfig(customLogger()))
	r.Use(gin.Recovery())
	// CORS設定
	r.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"PATCH",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			// "Authorization",
		},
		AllowOrigins: []string{
			"http://localhost",
		},
		MaxAge: 24 * time.Hour,
	}))
	defineRoutes(r)
	return r
}

func defineRoutes(r gin.IRouter) {
	v1 := r.Group("api/v1")
	// ヘルスチェック
	v1.GET("/healthcheck", healthcheck)
	// ユーザー
	{
		userController := controllers.UserController{}
		v1.GET("/users", userController.GetUserList)
		v1.GET("/users/:id", userController.GetUser)
		v1.POST("/users", userController.CreateUser)
	}
	// カテゴリ

	// アイテム
}

// 200OKを返すだけ
func healthcheck(c *gin.Context) {
	c.JSON(200, "OK")
}
