package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	AddTime  int    `json:"add_time"`
}

// TableName ...
func (User) TableName() string {
	return "user"
}

// InitQueuesTale 初始化queues表，如果不存在，则创建
func InitQueuesTale() {
	DB.AutoMigrate(&User{})
}
