package bootstrap

import (
	"gin-web/global"
	"gin-web/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
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
