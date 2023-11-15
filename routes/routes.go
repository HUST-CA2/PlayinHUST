package routes

import (
	"PlayinHUST/controller"
	"PlayinHUST/middleware"
	"PlayinHUST/view"

	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {

	//用户认证api
	authroute := r.Group("/auth")
	{
		authroute.POST("/register", controller.UserRegister)
		authroute.POST("/login", controller.UserLogin)
		authroute.POST("/submit", middleware.AuthMiddleware(), controller.UserSubmit)
		authroute.GET("/info", middleware.AuthMiddleware(), controller.UserInfo)
	}

	//展示页面
	pageroute := r.Group("/")
	{
		//首页
		pageroute.GET("/", view.IndexPage)
		pageroute.GET("/index", view.IndexPage)
		pageroute.GET("/index.html", view.IndexPage)

		//功能页
		pageroute.GET("/register", view.RegisterPage)
		pageroute.GET("/login", view.LoginPage)
		pageroute.GET("/submit", middleware.AuthMiddleware(), view.SubmitPage)
	}

	//管理静态文件的api，前端请求的img,js,css文件都放到"/src/*"
	srcroute := r.Group("/")
	{
		srcroute.Static("/src", "./public/resource")
	}

	return r
}
