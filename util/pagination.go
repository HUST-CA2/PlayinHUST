package util

import "gorm.io/gorm"

type Pagination struct {
	TotalItem int
	PageNow   int
	PageNum   int
	PageSize  int
}

// 在gorm的Scopes()函数内调用paginate(&pagination),需要先设置好pagination的参数
func Paginate(pagination *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pagination.PageNum = pagination.TotalItem / pagination.PageSize
		if pagination.TotalItem%pagination.PageSize != 0 {
			pagination.PageNum++
		}

		if pagination.PageNow <= 1 {
			pagination.PageNow = 1
		} else if pagination.PageNow >= pagination.PageNum {
			pagination.PageNow = pagination.PageNum
		}
		offset := (pagination.PageNow - 1) * pagination.PageSize
		return db.Limit(pagination.PageSize).Offset(offset)
	}

}
