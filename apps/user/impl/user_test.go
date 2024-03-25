package impl_test

import (
	"context"
	"testing"

	"github.com/liuli404/VBlog/apps/user"
	"github.com/liuli404/VBlog/apps/user/impl"
)

func init() {
	// 加载被测试对象 i 就是 User Server 的具体实现对象
	i = &impl.UserServiceImpl{}
}

var (
	i   user.Service
	ctx = context.Background()
)

func TestCreateUser(t *testing.T) {
	u, err := i.CreateUser(ctx, nil)
	// 单元测试异常处理，直接中断单元流程并且失败
	if err != nil {
		t.Fatal(err)
	}
	// 自定义测试异常
	if u == nil {
		t.Fatal("user not create")
	}
	// 正常打印
	t.Log(u)
}
