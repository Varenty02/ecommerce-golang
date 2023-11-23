package main

import (
	"Ecommerce/components/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteConfig(ctx appctx.AppContext, engine *gin.Engine) {
	//user route
	user := engine.Group("/user")
	user.GET("/profile", func(context *gin.Context) {
		context.JSON(http.StatusOK, "this is my profile")
	})
}
