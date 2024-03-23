package user

import "context"

// user.Service, 设计用户模块提供的接口

// 接口定义, 一定要考虑兼容性, 接口的参数不能变
type Service interface {
	// 创建用户
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// 查询用户列表
	QueryUser(context.Context, *QueryUserRequest) (*User, error)
	//查询用户详情
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
	//修改用户信息
	ModifyUser(context.Context, *ModifyUserRequest) (*User, error)
	// 删除用户
	DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
}

type CreateUserRequest struct {
	Username string
	Password string
	Role     string
	Label    map[string]string
}

type QueryUserRequest struct {
	PageSize   int
	PageNumber int
	Username   string
}

type DescribeUserRequest struct {
}

type ModifyUserRequest struct {
}

type DeleteUserRequest struct {
}
