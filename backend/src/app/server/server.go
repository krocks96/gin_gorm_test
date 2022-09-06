package server

import (
	"app/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// サーバ起動
func StartServer() {
	r := Router()
	r.Run()
}

func Router() *gin.Engine {
	r := gin.Default()
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
	// ルーティングの設定
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
