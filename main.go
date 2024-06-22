package main

import (
	"gin-web/bootstrap"
	"gin-web/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//初始化配置文件
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 启动服务器
	r.Run(":" + global.App.Config.App.Port)
}
