package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"image/logic"
	"strconv"
)

// AdjustBrightnessHandler 图片亮度调整
func AdjustBrightnessHandler(c *gin.Context)  {
	// 1. 解析参数
	imageID,_ := strconv.ParseInt(c.Param("id"), 10, 64)
	percentage,_ := strconv.ParseFloat(c.Query("percentage"), 64)
	// 2. logic层处理
	if err := logic.AdjustBrightness(imageID, percentage); err != nil {
		zap.L().Error("logic.AdjustBrightness failed", zap.Error(err))
		ResponseError(c, CodeSearchFailed)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, gin.H{
		"message": "success",
	})
}

// AdjustContrastHandler 图片对比度
func AdjustContrastHandler(c *gin.Context)  {
	// 1. 解析参数
	imageID,_ := strconv.ParseInt(c.Param("id"), 10, 64)
	percentage,_ := strconv.ParseFloat(c.Query("percentage"), 64)
	// 2. logic层处理
	if err := logic.AdjustContrast(imageID, percentage); err != nil {
		zap.L().Error("logic.AdjustContrast failed", zap.Error(err))
		ResponseError(c, CodeSearchFailed)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, gin.H{
		"message": "success",
	})
}

// AdjustGammaHandler gamma
func AdjustGammaHandler(c *gin.Context)  {
	// 1. 解析参数
	imageID,_ := strconv.ParseInt(c.Param("id"), 10, 64)
	gamma,_ := strconv.ParseFloat(c.Query("gamma"), 64)
	// 2. logic层处理
	if err := logic.AdjustGamma(imageID, gamma); err != nil {
		zap.L().Error("logic.AdjustGamma failed", zap.Error(err))
		ResponseError(c, CodeSearchFailed)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, gin.H{
		"message": "success",
	})
}

// AdjustSaturationHandler 图片饱和度
func AdjustSaturationHandler(c *gin.Context)  {
	// 1. 解析参数
	imageID,_ := strconv.ParseInt(c.Param("id"), 10, 64)
	percentage,_ := strconv.ParseFloat(c.Query("percentage"), 64)
	// 2. logic层处理
	if err := logic.AdjustSaturation(imageID, percentage); err != nil {
		zap.L().Error("logic.AdjustSaturation failed", zap.Error(err))
		ResponseError(c, CodeSearchFailed)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, gin.H{
		"message": "success",
	})
}

// BlurHandler 图片饱和度
func BlurHandler(c *gin.Context)  {
	// 1. 解析参数
	imageID,_ := strconv.ParseInt(c.Param("id"), 10, 64)
	sigma,_ := strconv.ParseFloat(c.Query("percentage"), 64)
	// 2. logic层处理
	if err := logic.Blur(imageID, sigma); err != nil {
		zap.L().Error("logic.Blur failed", zap.Error(err))
		ResponseError(c, CodeSearchFailed)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, gin.H{
		"message": "success",
	})
}

// CropCenterHandler 中心裁剪
func CropCenterHandler(c *gin.Context)  {
	// 1. 解析参数
	imageID,_ := strconv.ParseInt(c.Param("id"), 10, 64)
	width,_ := strconv.Atoi(c.Query("width"))
	height,_ := strconv.Atoi(c.Query("height"))
	// 2. logic层处理
	if err := logic.CropCenter(imageID, width, height); err != nil {
		zap.L().Error("logic.CropCenter failed", zap.Error(err))
		ResponseError(c, CodeSearchFailed)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, gin.H{
		"message": "success",
	})
}

// InvertHandler 中心裁剪
func InvertHandler(c *gin.Context)  {
	// 1. 解析参数
	imageID,_ := strconv.ParseInt(c.Param("id"), 10, 64)
	// 2. logic层处理
	if err := logic.Invert(imageID); err != nil {
		zap.L().Error("logic.Invert failed", zap.Error(err))
		ResponseError(c, CodeSearchFailed)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, gin.H{
		"message": "success",
	})
}