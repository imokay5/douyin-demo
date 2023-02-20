package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imokay5/douyin-demo/service"
)

type VideoData = service.VideoData

type FeedResponse struct {
	Response
	VideoList []VideoData `json:"video_list,omitempty"`
	NextTime  int64       `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
