package logic

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image/dao/mysql"
	"log"
)

func AdjustBrightness(imageID int64, percentage float64) (err error) {
	// 1. mysql获取图片路径
	path,err := mysql.GetImageByID(imageID)
	if err != nil{
		return err
	}
	fmt.Println(path)
	// 2. 打开文件 调整亮度
	// Open a test image.
	src, err := imaging.Open(path)
	if err != nil {
		log.Printf("failed to open image: %v", err)
	}
	img := imaging.AdjustBrightness(src, percentage)
	err = imaging.Save(img, "images/tmp/bbb.png")
	if err != nil {
		log.Printf("failed to save image: %v", err)
	}
	return
}


func AdjustContrast(imageID int64, percentage float64) (err error) {
	// 1. mysql获取图片路径
	path,err := mysql.GetImageByID(imageID)
	if err != nil{
		return err
	}
	fmt.Println(path)
	// 2. 打开文件 调整对比度
	src, err := imaging.Open(path)
	if err != nil {
		log.Printf("failed to open image: %v", err)
	}
	img := imaging.AdjustContrast(src, percentage)
	err = imaging.Save(img, "images/tmp/bbb.png")
	if err != nil {
		log.Printf("failed to save image: %v", err)
	}
	return
}

func AdjustGamma(imageID int64, gamma float64) (err error) {
	// 1. mysql获取图片路径
	path,err := mysql.GetImageByID(imageID)
	if err != nil{
		return err
	}
	fmt.Println(path)
	// 2. 打开文件 调整gamma
	src, err := imaging.Open(path)
	if err != nil {
		log.Printf("failed to open image: %v", err)
	}
	img := imaging.AdjustGamma(src, gamma)
	err = imaging.Save(img, "images/tmp/bbb.png")
	if err != nil {
		log.Printf("failed to save image: %v", err)
	}
	return
}

func AdjustSaturation(imageID int64, percentage float64) (err error) {
	// 1. mysql获取图片路径
	path,err := mysql.GetImageByID(imageID)
	if err != nil{
		return err
	}
	fmt.Println(path)
	// 2. 打开文件 调整饱和度
	src, err := imaging.Open(path)
	if err != nil {
		log.Printf("failed to open image: %v", err)
	}
	img := imaging.AdjustSaturation(src, percentage)
	err = imaging.Save(img, "images/tmp/bbb.png")
	if err != nil {
		log.Printf("failed to save image: %v", err)
	}
	return
}

func Blur(imageID int64, sigma float64) (err error) {
	// 1. mysql获取图片路径
	path,err := mysql.GetImageByID(imageID)
	if err != nil{
		return err
	}
	fmt.Println(path)
	// 2. 打开文件  高斯模糊
	src, err := imaging.Open(path)
	if err != nil {
		log.Printf("failed to open image: %v", err)
	}
	img := imaging.Blur(src, sigma)
	err = imaging.Save(img, "images/tmp/bbb.png")
	if err != nil {
		log.Printf("failed to save image: %v", err)
	}
	return
}

func CropCenter(imageID int64, width,height int) (err error) {
	// 1. mysql获取图片路径
	path,err := mysql.GetImageByID(imageID)
	if err != nil{
		return err
	}
	fmt.Println(path)
	// 2. 打开文件  中心裁剪
	src, err := imaging.Open(path)
	if err != nil {
		log.Printf("failed to open image: %v", err)
	}
	img := imaging.CropCenter(src, width, height)
	err = imaging.Save(img, "images/tmp/bbb.png")
	if err != nil {
		log.Printf("failed to save image: %v", err)
	}
	return
}

func Invert(imageID int64) (err error) {
	// 1. mysql获取图片路径
	path,err := mysql.GetImageByID(imageID)
	if err != nil{
		return err
	}
	fmt.Println(path)
	// 2. 打开文件  色系反转
	src, err := imaging.Open(path)
	if err != nil {
		log.Printf("failed to open image: %v", err)
	}
	img := imaging.Invert(src)
	err = imaging.Save(img, "images/tmp/bbb.png")
	if err != nil {
		log.Printf("failed to save image: %v", err)
	}
	return
}