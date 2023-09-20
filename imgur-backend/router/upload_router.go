package router

import (
	"imgur-backend/controller"
	"imgur-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	v1Group := r.Group("api")
	{
		// 上传图片
		v1Group.POST("/upload", controller.Upload)
	}

	r.Use(middleware.Cors())
	return r
}
