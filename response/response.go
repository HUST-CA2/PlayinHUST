package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 通用返回json的函数，http状态码和业务状态码code自定
func UniResp(ctx *gin.Context, httpStatus int, code int, msg string, data gin.H) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// 成功返回，http状态码为200，业务状态码为200
func SuccessResp(ctx *gin.Context, msg string, data gin.H) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": data,
	})
}

// 失败返回，http状态码为200，业务状态码为400
func FailedResp(ctx *gin.Context, msg string, data gin.H) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  msg,
		"data": data,
	})
}
