package common

import (
	"PlayinHUST/model"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func init() {
	//打开数据库连接
	var err error
	PlayinHUSTDB, err = gorm.Open(dbdriver, sqlConnection)
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	defer PlayinHUSTDB.Close()

	PlayinHUSTDB.AutoMigrate(&model.ClubInfo{})
	PlayinHUSTDB.AutoMigrate(&model.UserAccount{})

	sqlFile, err := ioutil.ReadFile("./common/dbinit.sql")
	if err != nil {
		log.Fatal(err)
	}

	PlayinHUSTDB.Exec(string(sqlFile))
}

// 把这个包init函数里的局部变量传到外面
func GetDB() *gorm.DB {
	return PlayinHUSTDB
}
