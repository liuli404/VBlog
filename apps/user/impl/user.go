package impl

import (
	"context"

	"github.com/liuli404/VBlog/apps/user"
)

// 判断服务有没有实现这个接口
// var _ user.Service = (*UserServiceImpl)(nil)

// 实现 User.Service 方法
// 创建用户
func (i *UserServiceImpl) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.User, error) {
	// 1. 校验用户请求参数
	if err := in.Validate(); err != nil {
		return nil, err
	}

	// 2. 创建用户实例对象
	u := user.NewUser(in)

	// 3. 把对象持久化（入库）
	// 定义orm与数据库的row关系
	if err := i.db.Create(u).Error; err != nil {
		return nil, err
	}

	// 4. 返回创建好的对象
	return u, nil
}

// 查询用户列表
func (i *UserServiceImpl) QueryUser(ctx context.Context, in *user.QueryUserRequest) (*user.User, error) {
	return nil, nil
}

// 查询用户详情
func (i *UserServiceImpl) DescribeUser(ctx context.Context, in *user.DescribeUserRequest) (*user.User, error) {
	return nil, nil
}
