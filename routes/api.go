package routes

import (
	"gin-web/app/controllers/app"
	"gin-web/app/middleware"
	"gin-web/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	setUserGroupRoutes(router)
}

func setUserGroupRoutes(router *gin.RouterGroup) {
	group := router.Group("/auth")
	group.POST("/register", app.Register)
	group.POST("/login", app.Login)
	authRouter := group.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.GET("/info", app.Info)
		authRouter.GET("/logout", app.Logout)
	}
}
