package controller

import (
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//to_user_id := c.Query("to_user_id")
	token := username + password

	//token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		// 关注
		// 连接数据库
		// 用户和目标用户follow和follower都改
		c.JSON(http.StatusOK, Response{StatusCode: 0})

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
