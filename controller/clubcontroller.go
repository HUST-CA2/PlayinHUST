package controller

import (
	"PlayinHUST/model"
	"PlayinHUST/response"
	"PlayinHUST/util"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 提交社团信
func SubmitClub(ctx *gin.Context) {
	clubname := ctx.PostForm("clubname")
	membergroup := ctx.PostForm("membergroup")
	clubinfo := ctx.PostForm("clubinfo")

	username, _ := ctx.Get("username")
	var accountName = username.(string)

	var club model.ClubInfo

	PlayinHUSTDB.Where("club_name = ? AND admin = ?", clubname, accountName).First(&club)
	if club.ID != 0 {
		response.FailedResp(ctx, "该社团信息已存在", gin.H{})
		return
	}

	club = model.ClubInfo{Admin: accountName, ClubName: clubname, MemberGroup: membergroup, ClubInfo: clubinfo}

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
		response.FailedResp(ctx, "删除失败", gin.H{})
	}
}

// 查询所有社团信息
func GetClubs(ctx *gin.Context) {
	//获取参数pagenow,pagesize,totalitem
	pagenow, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pagesize, _ := strconv.Atoi(ctx.DefaultQuery("pagesize", "5"))
	var totalitem int64 = 0
	PlayinHUSTDB.Model(&model.ClubInfo{}).Count(&totalitem)

	//参数赋值给pagination结构体
	pagination := util.Pagination{TotalItem: int(totalitem), PageNow: pagenow, PageSize: pagesize}

	//gorm scopes调用paginate函数，参数为pagination结构体
	//定义返回社团信息结构体，避免全部返回
	type clubInfosReturn struct {
		ID          uint
		Admin       string
		ClubName    string
		MemberGroup string
	}
	var clubinfosreturn []clubInfosReturn
	PlayinHUSTDB.Model(&model.ClubInfo{}).Scopes(util.Paginate(&pagination)).Find(&clubinfosreturn)

	response.SuccessResp(ctx, "查询信息成功", gin.H{
		"totalnum":  totalitem,
		"pagenum":   pagination.PageNum,
		"page":      pagination.PageNow,
		"pagesize":  pagesize,
		"clubinfos": clubinfosreturn,
	})
}

// 查询单个社团信息
func GetClub(ctx *gin.Context) {
	clubname := ctx.Param("clubname")
	clubinfo := model.ClubInfo{}

	PlayinHUSTDB.Where("club_name = ?", clubname).First(&clubinfo)
	if clubinfo.ID != 0 {
		response.SuccessResp(ctx, "查询信息成功", gin.H{
			"clubinfo": clubinfo,
		})
	} else {
		response.FailedResp(ctx, "查询失败，信息不存在", gin.H{})
	}
}
