package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"image/dao/mysql"
	"image/logic"
	"image/model"
)

func RegisterHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	var registerUser model.ParamRegisterUser
	if err := c.ShouldBindJSON(&registerUser); err != nil {
		zap.L().Error("register with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 业务处理
	if err := logic.Register(&registerUser); err != nil {
		zap.L().Error("logic.Register failed,", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExit) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// 1. 获取请求参数及参数校验
	p := new(model.User)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid params ", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed, ", zap.String("username:", p.UserName), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExit) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.UserName,
		"token":     user.AccessToken,
	})
}
