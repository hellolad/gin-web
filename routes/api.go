package routes

import (
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
	group := router.Group("/user")
	group.GET("/ping1", func(c *gin.Context) {
		c.String(http.StatusOK, "pong1")
	})
}
