package routes

import (
	v1 "LibraryManagementSys/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (b *BaseRouter) InitBaseRouter(baseRouter *gin.RouterGroup) {
	baseApi := v1.ApiGroupApp.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("register", baseApi.Register)
	}
}
