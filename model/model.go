package model

import "github.com/jinzhu/gorm"

type (
	//账户信息
	UserAccount struct {
		gorm.Model
		Account  string
		Password string
	}

	//数据库存的社团信息
	ClubInfo struct {
		gorm.Model
		Admin       string
		ClubName    string
		MemberGroup string
		ClubInfo    string
	}
)
