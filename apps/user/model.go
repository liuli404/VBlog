package user

import "time"

// 存放需要入库的数据

// 用户创建成功后返回的User对象
type User struct {
	Id       int   `gorm:"column:id"`
	CreateAt int64 `gorm:"column:created_at"`
	UpdateAt int64 `gorm:"column:updated_at"`
	// 用户参数
	*CreateUserRequest
}

// User的构造函数
func NewUser(req *CreateUserRequest) *User {
	return &User{
		CreateAt:          time.Now().Unix(),
		CreateUserRequest: req,
	}
}

// 定义对象存储的表的名称
func (u *User) TableName() string {
	return "users"
}
