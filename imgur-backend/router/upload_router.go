package router

import (
	"imgur-backend/controller"
	"imgur-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	v1Group := r.Group("api")
	{
		// 上传图片
		v1Group.POST("/upload", controller.Upload)
	}

	return r
}
