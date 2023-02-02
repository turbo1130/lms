package routes

import (
	"LibraryManagementSys/middleware"
	"github.com/gin-gonic/gin"
)

var ginServer *gin.Engine

func GinInit() {
	ginServer = gin.Default()
	ginServer.Use(middleware.JwtAuthHandler())
	userRoutes()
	ginServer.Run(":8081") // 必须放在最后
}

func userRoutes() {
	// 登录
	//ginServer.POST("/v/login", service.Login)

	/*	// 借阅
		groupLmsJyServer := ginServer.Group("/api/v1/jieyue")
		// 根据页面查询
		groupLmsJyServer.GET("/:page", service.JieyueDataByPage)
		// 所有出版社
		groupLmsJyServer.GET("/publishers", service.GetPublisherList)
		// 根据条件查询
		groupLmsJyServer.GET("/search/:page", service.JieyueDataByConditions)*/
}
