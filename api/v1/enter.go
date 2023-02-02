package v1

import "LibraryManagementSys/service"

type ApiGroup struct {
	UserApi
	BookApi
	BaseApi
	LendingRecordApi
	JwtInvalidApi
	RoleApi
}

var ApiGroupApp = new(ApiGroup) // 将所有的api都进行实例化注册

var (
	userService          = service.GroupsService.UserService
	lendingRecordService = service.GroupsService.LendingRecordService
	bookService          = service.GroupsService.BookService
	roleService          = service.GroupsService.RoleService
)
