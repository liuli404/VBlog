package impl

import (
	"context"

	"github.com/liuli404/VBlog/apps/user"
)

// 判断服务有没有实现这个接口
// var _ user.Service = (*UserServiceImpl)(nil)

// 实现 User.Service 方法
func (I *UserServiceImpl) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.User, error) {
	return nil, nil
}

// 查询用户列表
func (I *UserServiceImpl) QueryUser(ctx context.Context, in *user.QueryUserRequest) (*user.User, error) {
	return nil, nil
}

// 查询用户详情
func (I *UserServiceImpl) DescribeUser(ctx context.Context, in *user.DescribeUserRequest) (*user.User, error) {
	return nil, nil
}

// 修改用户信息
func (I *UserServiceImpl) ModifyUser(ctx context.Context, in *user.ModifyUserRequest) (*user.User, error) {
	return nil, nil
}

// 删除用户
func (I *UserServiceImpl) DeleteUser(ctx context.Context, in *user.DeleteUserRequest) (*user.User, error) {
	return nil, nil
}
