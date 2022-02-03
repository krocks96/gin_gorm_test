package server

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "app/config"
    "app/controllers"
)

func StartServer() error {
    fmt.Println("Rest API with Mux Routers")
    router := gin.Default()
    defineRoutes(router)
    return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.ServerPort), router)
}

func defineRoutes(r gin.IRouter) {
    v1 := r.Group("/v1")
    // ユーザー
    {
        userController := controllers.UserController{}
        v1.GET("/users", userController.GetAllUser)
        v1.GET("/users/:id", userController.GetUser)
        v1.POST("/users", userController.CreateUser)
    }
}