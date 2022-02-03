package controllers

import (
    "net/http"
	"log"
	"strconv"
    "github.com/gin-gonic/gin"
	"app/models"
	"app/models/entity"
)

type UserController struct {}

func (pc UserController) GetAllUser(c *gin.Context){
	// ユーザー取得
	users, err := model.FindAllUser()
	if err != nil {
		c.AbortWithStatus(400)
		log.Println(err)
	} else {
		c.JSONP(http.StatusOK, gin.H{
			"message": "ok",
			"data": users,
		})
	}
}

func (pc UserController) GetUser(c *gin.Context){
	// バリデーションは要検討
	strId := c.Param("id")
	if strId == "" {
		c.AbortWithStatus(400)
		log.Println("idを入力してください")
	}
	id, _ := strconv.Atoi(strId)
	// ユーザー取得
	user, err := model.FindUser(id)
	if err != nil {
		c.AbortWithStatus(400)
		log.Println(err)
	} else {
		c.JSONP(http.StatusOK, gin.H{
			"message": "ok",
			"data": user,
		})
	}
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
