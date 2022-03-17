package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code    MyCode      `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseErrorWithMsg(ctx *gin.Context, code MyCode, msg interface{}) {
	ctx.JSON(http.StatusOK, &ResponseData{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

func ResponseError(ctx *gin.Context, code MyCode) {
	ctx.JSON(http.StatusInternalServerError, &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	})
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	})
}
