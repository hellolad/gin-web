package bootstrap

import (
	"gin-web/app/middleware"
	"gin-web/global"
	"gin-web/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	if global.App.Config.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	// 默认
	//router := gin.Default()
	router := gin.New()
	router.Use(gin.Logger(), middleware.CustomRecovery())

	// 跨域处理
	router.Use(middleware.Cors())

	router.StaticFile("/", "./static/index.html")

	group := router.Group("/api")
	routes.SetApiGroupRoutes(group)
	return router
}

func RunServer() {
	// 启动服务器
	r := setupRouter()
	err := r.Run(":" + global.App.Config.App.Port)
	if err != nil {
		return
	}
}
