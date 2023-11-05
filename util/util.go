package util

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"
)

// GetCurrentFormatStr 获取标准时间字符串
func GetCurrentFormatStr(fmtStr string) string {
	if fmtStr == "" {
		fmtStr = "2006-01-02 15:04:05"
	}
	return time.Now().Format(fmtStr)
}

// MD5 用于对用户密码加密
func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

// getCookieUsername 获取cookie值中的用户名
func GetCookieUsername(cookieValue string) string {
	nameidx := strings.LastIndex(cookieValue, ":")
	accountName := cookieValue[0:nameidx]
	return accountName

}
