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
	v1.POST("/register", controller.RegisterHandler) // 注册
	v1.POST("/login", controller.LoginHandler)       // 登录

	v1.GET("/image/bright/:id", controller.AdjustBrightnessHandler) // 图片亮度
	v1.GET("/image/contrast/:id", controller.AdjustContrastHandler) // 图片对比度
	v1.GET("/image/gamma/:id", controller.AdjustGammaHandler) // Gamma
	v1.GET("/image/saturation/:id", controller.AdjustSaturationHandler) // 图片饱和度
	v1.GET("/image/blur/:id", controller.BlurHandler) // 高斯模糊
	v1.GET("/image/cropCenter/:id", controller.CropCenterHandler) // 中心裁剪
	v1.GET("/image/invert/:id", controller.InvertHandler) // 色系反转


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
