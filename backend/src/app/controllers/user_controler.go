package controllers

import (
    "net/http"
	"log"
    "github.com/gin-gonic/gin"
	"app/models"
	"app/models/entity"
)

type UserController struct {}

func (pc UserController) GetAllUser(c *gin.Context){
    c.JSONP(http.StatusOK, gin.H{
        "message": "ok",
        "data": "AllUsers",
    })
}

func (pc UserController) GetUser(c *gin.Context){
	id := c.Param("id")
    c.JSONP(http.StatusOK, gin.H{
        "message": "ok",
        "data": "User" + id,
    })
}

func (pc UserController) CreateUser(c *gin.Context){
	var req entity.User
	// エラー
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(400)
		log.Println(err)
	}
	// ユーザー作成
	user := entity.User{Name: req.Name}
	_, err := model.CreateUser(&user)
	if err != nil {
		c.AbortWithStatus(400)
		log.Println(err)
	} else {
		c.JSON(204, nil)
	}
}
