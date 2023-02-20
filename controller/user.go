package controller

import (
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imokay5/douyin-demo/repository"
	"github.com/imokay5/douyin-demo/service"
)

type User = repository.User

// 登录信息 token：账号密码
var usersLoginInfo = map[string]User{
	"xch123456": {
		Id:         1001,
		Name:       "xch",
		Password:   "123456",
		CreateTime: time.Now(),
		ModifyTime: time.Now(),
	},
}

var userIdSequence = int64(1)

//请求用户信息的返回
type UserResponse struct {
	Response
	User service.UserData `json:"user,omitempty"`
}

//登陆或者注册结束后返回给客户端
type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

// 获取用户信息
func UserInfo(c *gin.Context) {
	user_id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64) // 将查询 字符串 转为 10进制数据，限制8位
	// 可以通过Query来获取URL中？后面所携带的参数。
	token := c.Query("token")
	userData, err := service.QueryUserData(user_id, token)
	if err == nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     *userData,
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
