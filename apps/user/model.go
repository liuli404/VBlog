package user

// 存放需要入库的数据

// 用户创建成功后返回的User对象
type User struct {
	Id       int
	CreateAt int64
	UpdateAt int64
	// 用户参数
	*CreateUserRequest
}
