package controller

import (
	"PlayinHUST/common"
	"PlayinHUST/model"
	"PlayinHUST/util"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var PlayinHUSTDB *gorm.DB = common.GetDB()

func UserRegister(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")
	confirm := ctx.PostForm("confirm")

	//数据验证
	if len(account) == 0 {
		ctx.HTML(200, "register.html", gin.H{
			"title":   "注册",
			"head":    "注册",
			"warning": "账户不能为空",
		})
	} else if len(password) < 6 {
		ctx.HTML(200, "register.html", gin.H{
			"title":   "注册",
			"head":    "注册",
			"warning": "密码不能小于6位",
		})
	} else if password != confirm {
		ctx.HTML(200, "register.html", gin.H{
			"title":   "注册",
			"head":    "注册",
			"warning": "两次密码不一致，请重新输入",
		})
	} else {
		oldAccount := model.UserAccount{}
		PlayinHUSTDB.Where("account = ?", account).First(&oldAccount)
		if oldAccount.ID != 0 {
			ctx.HTML(200, "register.html", gin.H{
				"title":   "注册",
				"head":    "注册",
				"warning": "账户已存在，请更换账户名",
			})

		} else {
			newAccount := model.UserAccount{Account: account, Password: util.MD5(password)}

			err := PlayinHUSTDB.Save(&newAccount)
			if err != nil {
				fmt.Println(err.Error)
			}
			ctx.HTML(200, "login.html", gin.H{
				"title":   "登录",
				"head":    "登录",
				"welcome": "注册成功，请",
			})
		}
	}
}

func UserLogin(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")

	//数据验证
	if len(account) == 0 {
		ctx.HTML(200, "login.html", gin.H{
			"title":   "登录",
			"head":    "登录",
			"warning": "账户不能为空",
		})
	} else if len(password) == 0 {
		ctx.HTML(200, "login.html", gin.H{
			"title":   "登录",
			"head":    "登录",
			"warning": "密码不能为空",
		})
	} else {
		oldAccount := model.UserAccount{}
		PlayinHUSTDB.Where("account = ?", account).First(&oldAccount)
		if oldAccount.ID == 0 {
			ctx.HTML(200, "login.html", gin.H{
				"title":   "登录",
				"head":    "登录",
				"warning": "账户不存在，请先注册",
			})
		} else if util.MD5(password) != oldAccount.Password {
			ctx.HTML(200, "login.html", gin.H{
				"title":   "登录",
				"head":    "登录",
				"warning": "密码错误",
			})
		} else {
			//发放token
			token, err := common.ReleaseToken(oldAccount)
			if err != nil {
				ctx.HTML(200, "login.html", gin.H{
					"title":   "登录",
					"head":    "登录",
					"warning": "系统异常",
				})
				log.Printf("token generate error: %v", err)
				return
			}

			ctx.HTML(200, "test.html", gin.H{"token": token})
		}
	}

}

func UserSubmit(ctx *gin.Context) {
	clubname := ctx.PostForm("clubname")
	membergroup := ctx.PostForm("qqgroup")
	clubinfo := ctx.PostForm("clubinfo")

	username, _ := ctx.Get("username")
	var accountName = username.(string)

	club := model.ClubInfo{Admin: accountName, ClubName: clubname, MemberGroup: membergroup, ClubInfo: clubinfo}

	err := PlayinHUSTDB.Save(&club)
	if err != nil {
		fmt.Println(err.Error)
	}

	ctx.HTML(200, "congratulation.html", gin.H{
		"account":     accountName,
		"name":        clubname,
		"membergroup": membergroup,
		"ClubInfo":    clubinfo,
	})

}

func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userAccount := user.(model.UserAccount)
	userdto := model.ToUserDto(userAccount)

	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取信息成功",
		"data": gin.H{"user": userdto},
	})
}
