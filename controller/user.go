package controller

import (
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

// 登录信息 token：账号密码
var usersLoginInfo = map[string]User{
	"zhangleidouyin": { // token
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserResponse struct {
	Response
	User User `json:"user,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

// 获取用户信息
func UserInfo(c *gin.Context) {
	// 可以通过Query来获取URL中？后面所携带的参数。
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "用户不存在"},
		})
	}
}

// 注册
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	token := username + password
	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "用户已经创建"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1) // 用户 id 序列加一
		newUser := User{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    token,
		})
	}
}

// 登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "用户不存在，请注册"},
		})
	}
}
