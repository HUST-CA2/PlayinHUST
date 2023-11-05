package middleware

import (
	"PlayinHUST/common"
	"PlayinHUST/model"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取Authorization header
		tokenString := ctx.GetHeader("Authorization")

		//验证格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "Token格式有问题",
				"data": "",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "Token无效",
				"data": "",
			})
			ctx.Abort()
			return
		}

		//通过验证后获取claims里面的userid
		userId := claims.UserId
		DB := common.GetDB()
		var user model.UserAccount
		DB.First(&user, userId)

		//若用户不存在
		if user.ID == 0 {
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  "用户不存在",
				"data": "",
			})
			ctx.Abort()
		}

		//若用户存在，将名字以字符串形式存入上下文
		ctx.Set("username", user.Account)
		//若用户存在，将信息以json形式存入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
