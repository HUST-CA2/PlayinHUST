package controller

import (
	"PlayinHUST/common"
	"PlayinHUST/model"
	"PlayinHUST/response"
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
		response.FailedResp(ctx, "账户不能为空", gin.H{})
	} else if len(password) < 6 {
		response.FailedResp(ctx, "密码不能小于6位", gin.H{})
	} else if password != confirm {
		response.FailedResp(ctx, "两次密码不一致", gin.H{})
	} else {
		oldAccount := model.UserAccount{}
		PlayinHUSTDB.Where("account = ?", account).First(&oldAccount)
		if oldAccount.ID != 0 {

			response.FailedResp(ctx, "账户已存在，请更换用户名", gin.H{})

		} else {
			newAccount := model.UserAccount{Account: account, Password: util.MD5(password)}

			err := PlayinHUSTDB.Save(&newAccount)
			if err != nil {
				fmt.Println(err.Error)
			}

			response.SuccessResp(ctx, "注册成功", gin.H{})
		}
	}
}

func UserLogin(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")

	//数据验证
	if len(account) == 0 {
		response.FailedResp(ctx, "账户不能为空", gin.H{})
	} else if len(password) == 0 {
		response.FailedResp(ctx, "密码不能为空", gin.H{})
	} else {
		oldAccount := model.UserAccount{}
		PlayinHUSTDB.Where("account = ?", account).First(&oldAccount)
		if oldAccount.ID == 0 {

			response.FailedResp(ctx, "账户不存在，请先注册", gin.H{})
		} else if util.MD5(password) != oldAccount.Password {

			response.FailedResp(ctx, "密码错误", gin.H{})
		} else {
			//发放token
			token, err := common.ReleaseToken(oldAccount)
			if err != nil {
				response.FailedResp(ctx, "发放token异常", gin.H{})
				log.Printf("token generate error: %v", err)
				return
			}

			response.SuccessResp(ctx, "获取token成功", gin.H{"token": token})
		}
	}

}

func UserSubmit(ctx *gin.Context) {
	clubname := ctx.PostForm("clubname")
	membergroup := ctx.PostForm("membergroup")
	clubinfo := ctx.PostForm("clubinfo")

	username, _ := ctx.Get("username")
	var accountName = username.(string)

	club := model.ClubInfo{Admin: accountName, ClubName: clubname, MemberGroup: membergroup, ClubInfo: clubinfo}

	err := PlayinHUSTDB.Save(&club)
	if err != nil {
		fmt.Println(err.Error)
	}

	response.SuccessResp(ctx, "提交社团信息成功", gin.H{
		"account":     accountName,
		"clubname":    clubname,
		"membergroup": membergroup,
		"clubinfo":    clubinfo,
	})
}

func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userAccount := user.(model.UserAccount)
	userdto := model.ToUserDto(userAccount)

	response.SuccessResp(ctx, "获取用户信息成功", gin.H{
		"user": userdto,
	})
}
