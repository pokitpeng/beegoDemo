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
