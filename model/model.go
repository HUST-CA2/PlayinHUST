package model

import "github.com/jinzhu/gorm"

type (
	//账户信息
	UserAccount struct {
		gorm.Model
		Account  string
		Password string
	}

	//返回给前端的信息
	UserDto struct {
		Account string
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

func ToUserDto(user UserAccount) UserDto {
	return UserDto{
		Account: user.Account,
	}
}
