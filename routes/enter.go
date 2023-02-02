package routes

type RouterGroup struct {
	BaseRouter
	BookRouter
	LendingRecordRouter
	RoleRouter
	UserRouter
	JwtInvalidRouter
}

var RouterGroups = new(RouterGroup) // 将所有的router都进行实例化注册
