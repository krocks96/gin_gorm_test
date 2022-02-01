package main

import (
        "app/server"
        "app/db"
        )

func main() {
    // DB起動
    db.Init()
    // サーバー起動
    server.StartServer()
    // DBクローズ
    db.Close()

    // DB処理
    // dsn := "docker:root@root(127.0.0.1:3306)/backend?charset=utf8mb4&parseTime=True"
    // db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    // if err != nil {
    //   panic("failed to connect database")
    // }    

    // // ルーティング
    // router := gin.Default()
    // userController := controller.UserController{}
    // router.GET("/users", userController.GetAllUser)
    // router.GET("/users/:id", userController.GetUser)

    // router.Run(":3000")
}