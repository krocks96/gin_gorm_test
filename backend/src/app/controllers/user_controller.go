package controllers

import (
	model "app/models"
	"app/parameters"
	"app/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (pc UserController) GetUserList(c *gin.Context) {
	// ユーザー取得
	users, err := model.FindAllUser()
	if err != nil {
		response.Render400(c, err)
		return
	}
	response.RenderUserList(c, users)
}

func (pc UserController) GetUser(c *gin.Context) {
	// パスパラメータ
	strId := c.Param("id")
	Id, err := strconv.ParseUint(strId, 10, 64)
	if Id == 0 {
		response.Render400(c, "Invalid Id")
		return
	}
	// ユーザー取得
	user, err := model.FindUser(Id)
	if err != nil {
		response.Render400(c, err)
		return
	}
	response.RenderUser(c, user)
}

func (pc UserController) CreateUser(c *gin.Context) {
	// リクエストパラメータのバインド
	var req parameters.PostUserParams
	if err := c.Bind(&req); err != nil {
		response.Render400(c, err)
		return
	}
	// ユーザー作成
	_, err := model.CreateUser(req)
	if err != nil {
		response.Render400(c, err)
		return
	}
	response.Render201(c)
}
