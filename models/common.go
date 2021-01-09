package models

import "gorm.io/gorm"

var DB *gorm.DB

// InitTales 初始化所有表，如果不存在，则创建
func InitTales() {
	_ = DB.AutoMigrate(&User{})
	_ = DB.AutoMigrate(&Author{})
	_ = DB.AutoMigrate(&Book{})
}
