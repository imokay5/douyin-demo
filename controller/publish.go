package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imokay5/douyin-demo/repository"
	"github.com/imokay5/douyin-demo/service"
)

// 视频列表响应
type VideoListResponse struct {
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
	saveFile := filepath.Join("./public/videos/", finalName)
	fmt.Println("saveFile:", saveFile)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 根据 文件名，文件标题，用户id 上传视频
	newVid, err := service.PostVideo(finalName, title, user.Id)
	if newVid == nil || err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 2,
			StatusMsg:  "save post video to db fail",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// 获取当前登录用户发布的视频列表
func PublishList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	videoList, err := service.GetPublishList(userId, token)
	if err != nil {
		fmt.Printf("get publish list failed: %s", err)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "get publish list failed",
			},
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: videoList,
		})
	}
}
