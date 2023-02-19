package controller

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/imokay5/douyin-demo/repository"
)

// 视频列表响应
type VideoListRespanse struct {
	Response              // 响应
	VideoList []VideoData `json:"video_data"` // 视频列表
}

// 发布检查 token，将上传视频保存于 public 目录，将数据保存于 db 中
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	fmt.Println("---", token, "---", title)
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": data,
		})
	}
	filename := filepath.Base(data.Filename)
	user := repository.UsersLoginInfo[token]
	fmt.Println("user.Password:", user.Password)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public", finalName)
	fmt.Println("saveFile:", saveFile)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
}
