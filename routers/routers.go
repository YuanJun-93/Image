package routers

import (
	"github.com/gin-gonic/gin"
	"image/controller"
	"image/logger"
	"image/middlewares"
	"net/http"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()
	//r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	// 用户功能
	/*
		eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2NTQyODkwMzk4Mzg0NTM3NiwidXNlcm5hbWUiOiJ5dWFuanVuIiwiZXhwIjoxNjUwMDI4NTY0LCJpc3MiOiJpbWFnZSJ9.MtkB_93EE4iwuqvhKTBAAzvXpRHwoUOg2EN_yPwvJB4
	*/
	v1.POST("/register", controller.RegisterHandler) // 注册
	v1.POST("/login", controller.LoginHandler)       // 登录

	// 需要登录之后才能做的操作
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		// Token testing
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

		// images operator
		v1.GET("/store", controller.GetImages)            // 获取图片
		v1.POST("/store/upload", controller.UploadImages) // 上传图片
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
