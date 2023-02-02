package routes

import (
	v1 "LibraryManagementSys/api/v1"
	"github.com/gin-gonic/gin"
)

type JwtInvalidRouter struct {
}

func (j *JwtInvalidRouter) InitJwtInvalidRouter(Router *gin.RouterGroup) {
	jwtInvalidApi := v1.ApiGroupApp.JwtInvalidApi
	Router.POST("logout", jwtInvalidApi.Logout)
}
