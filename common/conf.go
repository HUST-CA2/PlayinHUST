package common

import "github.com/jinzhu/gorm"

//这个文件用来配置

//服务器配置
var (
	ServerName string = "localhost"
	Port       string = ":8080"
)

//数据库配置
var (
	dbdriver      = "mysql"
	dbusername    = "root"
	dbpassword    = "Guo20040607"
	dbaddr        = "(localhost:3306)"
	dbname        = "playinhustdb"
	dboption      = "multiStatements=true&&parseTime=true"
	PlayinHUSTDB  *gorm.DB
	sqlConnection = dbusername + ":" + dbpassword + "@tcp" + dbaddr + "/" + dbname + "?" + dboption
)
