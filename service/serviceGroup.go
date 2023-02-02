package service

type GroupService struct {
	UserService
	LendingRecordService
	BookService
	RoleService
}

var GroupsService = new(GroupService)
