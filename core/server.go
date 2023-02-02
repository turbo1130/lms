package core

import (
	"LibraryManagementSys/docs"
	"LibraryManagementSys/global"
	"LibraryManagementSys/middleware"
	"LibraryManagementSys/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunGinServer() {
	Router := gin.Default()
	routerGroups := routes.RouterGroups
	// swagger
	docs.SwaggerInfo.BasePath = global.LMS_CONFIG.System.RouterPrefix
	Router.GET(global.LMS_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LMS_LOG.Info("register swagger handler")

	// 注册路由
	privateGroup := Router.Group(global.LMS_CONFIG.System.RouterPrefix)
	routerGroups.BaseRouter.InitBaseRouter(privateGroup)
	// 以下请求需要拦截
	privateGroup.Use(middleware.JwtAuthHandler())
	{
		routerGroups.JwtInvalidRouter.InitJwtInvalidRouter(privateGroup)
		routerGroups.BookRouter.InitBookRouter(privateGroup)
		routerGroups.LendingRecordRouter.InitLendingRecordRouter(privateGroup)
		routerGroups.UserRouter.InitUserRouter(privateGroup)
		routerGroups.RoleRouter.InitRoleRouter(privateGroup)
	}

	_ = Router.Run(":" + global.LMS_CONFIG.Gin.Port) // 必须放在最后
}
