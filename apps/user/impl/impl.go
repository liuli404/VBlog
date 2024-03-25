package impl

import (
	"github.com/liuli404/VBlog/conf"
	"gorm.io/gorm"
)

// 定义一个结构体，实现接口
type UserServiceImpl struct {
	// 依赖了一个数据库操作的连接池对象
	db *gorm.DB
}

// 获取全局DB对象
func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		db: conf.C().DB(),
	}
}
