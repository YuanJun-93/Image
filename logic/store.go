package logic

import (
	"image/dao/mysql"
	"image/model"
)

func GetImages() (imageInfo []*model.Image, err error) {
	// 1. 从数据库获取
	return mysql.GetImages()
}

func UploadImages(ImageInfo *model.Image) (err error) {
	// 1. 插入到数据库
	return mysql.UploadImages(ImageInfo)
}
