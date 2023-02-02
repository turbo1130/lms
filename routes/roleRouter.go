package routes

import (
	v1 "LibraryManagementSys/api/v1"
	"github.com/gin-gonic/gin"
)

type RoleRouter struct {
}

func (r *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	group := Router.Group("role")
	roleApi := v1.ApiGroupApp.RoleApi
	{
		group.PUT("ur", roleApi.UpdateUserRole)
		group.GET("list", roleApi.GetRoleList)
	}
}
