package router

import (
	"github.com/bricsport/common/exceptions"
	"github.com/bricsport/handler"
	"github.com/gin-gonic/gin"
)

var prefix = "/api/v2"

func LoadUserRoutes(router *gin.Engine, controller handler.UserHandler) {

	authRouter := router.Group(prefix)

	// router.GET("/")

	authRouter.POST("/", makeGinFunc(controller.Create))

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}

// TODO: move this type later
type controller func(ctx *gin.Context) (*handler.Payload, *exceptions.Exception)

func makeGinFunc(fn controller) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := fn(ctx)
		if err != nil {
			ctx.JSON(err.Code, err)
		}

		ctx.JSON(resp.Code, resp)
	}
}
