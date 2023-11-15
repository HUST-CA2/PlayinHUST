package view

import (
	"github.com/gin-gonic/gin"
)

func IndexPage(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{})
}

func RegisterPage(ctx *gin.Context) {
	ctx.HTML(200, "register.html", gin.H{
		"title": "注册",
		"head":  "注册",
	})
}

func LoginPage(ctx *gin.Context) {
	ctx.HTML(200, "login.html", gin.H{
		"title": "登录",
		"head":  "登录",
	})
}

func SubmitPage(ctx *gin.Context) {

	username, _ := ctx.Get("username")

	var accountName = username.(string)
	ctx.HTML(200, "submit.html", gin.H{
		"account": accountName,
		"title":   "提交",
		"head":    "提交",
	})
}
