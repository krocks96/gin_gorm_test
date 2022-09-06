package response

import (
	"app/models/entity"

	"log"

	"github.com/gin-gonic/gin"
)

type responseUserList struct {
	UserList []user `json:"userList"`
}

func Render200(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, data)
}

func Render201(ctx *gin.Context) {
	ctx.JSON(201, nil)
}

func Render204(ctx *gin.Context) {
	ctx.JSON(204, nil)
}

func Render400(ctx *gin.Context, err interface{}) {
	var res customError
	switch v := err.(type) {
	case error:
		res.Message = v.Error()
	case string:
		res.Message = v
	default:
		res.Message = "Unexpected Error. Please check log."
	}
	log.Println(err)
	ctx.JSON(400, res)
}

// ユーザー関連
func generateUserSchema(u *entity.User) user {
	return user{
		ID:          u.ID,
		DisplayName: u.DisplayName,
	}
}

func RenderUserList(c *gin.Context, data []*entity.User) {
	var res responseUserList
	res.UserList = []user{}
	for _, v := range data {
		u := generateUserSchema(v)
		res.UserList = append(res.UserList, u)
	}
	Render200(c, res)
}

func RenderUser(c *gin.Context, data *entity.User) {
	res := generateUserSchema(data)
	Render200(c, res)
}
