package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"image/logic"
	"image/model"
	"image/pkg/snowflake"
)

func GetImages(c *gin.Context) {
	// 1. logic层获取store的image
	data, err := logic.GetImages()
	if err != nil {
		zap.L().Error("logic get images failed", zap.Error(err))
		ResponseError(c, CodeSearchFailed)
		return
	}
	ResponseSuccess(c, data)
}

func UploadImages(c *gin.Context) {
	// 1. 判断参数是否存在
	upload := new(model.ParamImagesUpload)
	if err := c.ShouldBindJSON(upload); err != nil {
		zap.L().Error("param upload failed", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 2. logic 生成image信息，传入db
	image_id := snowflake.GenID()
	user_id, err := GetCurrentUserID(c)
	if err != nil {
		zap.L().Error("user not login", zap.Error(err))
		return
	}
	ImageInfo := model.Image{
		ImageID:   image_id,
		UserID:    user_id,
		ImageCost: upload.ImageCost,
		ImageName: upload.ImageName,
		ImageUrl:  upload.ImageUrl,
	}
	if err := logic.UploadImages(&ImageInfo); err != nil {
		zap.L().Error("logic upload images faild", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
