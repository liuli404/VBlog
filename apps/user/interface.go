package user

import (
	"context"

	"github.com/go-playground/validator/v10"
)

// user.Service, 设计用户模块提供的接口

// validator.New()校验器对象
var v = validator.New()

// 接口定义, 一定要考虑兼容性, 接口的参数不能变
type Service interface {
	// 创建用户
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// 查询用户列表
	QueryUser(context.Context, *QueryUserRequest) (*User, error)
	//查询用户详情
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
}

type CreateUserRequest struct {
	Username string            `validate:"required" gorm:"column:username"`
	Password string            `validate:"required" gorm:"column:password"`
	Role     string            `validate:"required" gorm:"column:role"`
	Label    map[string]string `gorm:"column:label;serializer:json"`
}

// 提供一个构造函数，防止对象内部出现空指针，指针对象为初始化。
// 还可以提供默认值
func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Role:  "member",
		Label: map[string]string{},
	}
}

// 参数校验方法
func (req *CreateUserRequest) Validate() error {
	// 1. 自定义报错
	// if req.Username == "" {
	// 	return fmt.Errorf("Username required")
	// }
	// if req.Password == "" {
	// 	return fmt.Errorf("Password required")
	// }

	// 2. validator 库
	return v.Struct(req)
}

type QueryUserRequest struct {
	PageSize   int
	PageNumber int
	Username   string
}

type DescribeUserRequest struct {
}
