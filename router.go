package main

// xch
import (
	"github.com/gin-gonic/gin"
	"github.com/imokay5/douyin-demo/controller"
)

/*初始化路由*/
func initRouter(r *gin.Engine) {
	r.Static("/static", "./public") // Public目录用于提供静态资源
	apiRouter := r.Group("/douyin")

	// basic apis
	// POST: 127.0.0.1:8080/douyin/publish/action/		token=xch123456&title=a_little_bear
	apiRouter.POST("publish/action/", controller.Publish) // 发布视频

}
