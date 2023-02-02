package routes

import (
	v1 "LibraryManagementSys/api/v1"
	"LibraryManagementSys/constant"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(userRouter *gin.RouterGroup) {
	group := userRouter.Group("user/self")
	userApi := v1.ApiGroupApp.UserApi
	{
		group.GET(constant.EMPTY, userApi.GetUserInfo)
		group.PUT(constant.EMPTY, userApi.UpdateUserInfo)
		group.POST(constant.EMPTY, userApi.ChangeUserPassword)
	}
}
