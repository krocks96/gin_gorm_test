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

	userController := controllers.UserController{}
	router.GET("/users", userController.GetAllUser)
    router.GET("/users/:id", userController.GetUser)
	router.POST("/users", userController.CreateUser)
    return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.ServerPort), router)
}