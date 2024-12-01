package router

import (
	"golang-pos/controller"

	"github.com/gin-gonic/gin"
)

func Routes(app *gin.Engine) {
	route := app
	route.GET("/user", controller.UsersGet)

	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

}
