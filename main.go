package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/imokay5/douyin-demo/repository"
)

func main() {
	// 初始化数据库
	if err := repository.Init(); err != nil {
		// Exit 函数可以让当前程序以给出的状态码 code 退出。一般来说，状态码 0 表示成功，非 0 表示出错。
		os.Exit(-1)
	}

	r := gin.Default()
	initRouter(r)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()
}
