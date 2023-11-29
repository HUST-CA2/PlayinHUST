package controller

import (
	"PlayinHUST/model"
	"PlayinHUST/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 提交社团信息
func SubmitClub(ctx *gin.Context) {
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
		"admin":       accountName,
		"clubname":    clubname,
		"membergroup": membergroup,
		"clubinfo":    clubinfo,
	})
}

// 删除社团信息
func DeleteClub(ctx *gin.Context) {
	clubname := ctx.Param("clubname")
	username, _ := ctx.Get("username")
	var accountName = username.(string) //accountName是从当前token获取的用户名

	clubinfo := model.ClubInfo{}

	PlayinHUSTDB.Where("club_name = ? AND admin = ?", clubname, accountName).First(&clubinfo)
	if clubinfo.ID != 0 {
		PlayinHUSTDB.Delete(&clubinfo)
		response.SuccessResp(ctx, "删除社团信息成功", gin.H{
			"admin":       accountName,
			"clubname":    clubname,
			"membergroup": clubinfo.MemberGroup,
			"clubinfo":    clubinfo.ClubInfo,
		})
	} else {
		response.FailedResp(ctx, "没有权限", gin.H{})
	}
}
