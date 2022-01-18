package http

import (
	"github.com/gin-gonic/gin"
	userController "github.com/wa1kman999/goblog-web-svc/internal/controller/user"
	"net/http"
)

const (
	v1prefix = "/goblog/v1"
)

func InitRouter(router *gin.Engine) {

	router.GET("/ready", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	router.GET("/healthy", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	// 用户模块
	user := router.Group(v1prefix + "/user")
	{
		// 登陆页面
		user.POST("login", userController.Login)
	}

}
