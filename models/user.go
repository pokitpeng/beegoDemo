package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}

// TableName ...
func (User) TableName() string {
	return "user"
}

// InitUserTale 初始化user表，如果不存在，则创建
func InitUserTale() {
	_ = DB.AutoMigrate(&User{})
}
