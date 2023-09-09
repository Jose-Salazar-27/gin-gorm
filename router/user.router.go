package router

import "github.com/gin-gonic/gin"

var prefix = "/api/v2"

func LoadUserRoutes(router *gin.Engine) {

	// gin.HandlerFunc
	authRouter := router.Group(prefix)

	authRouter.GET("/ping", func(c *gin.Context) {
		c.String(200, "auth pong")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
