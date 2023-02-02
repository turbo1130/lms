package routes

import (
	v1 "LibraryManagementSys/api/v1"
	"LibraryManagementSys/constant"
	"github.com/gin-gonic/gin"
)

type LendingRecordRouter struct {
}

func (l *LendingRecordRouter) InitLendingRecordRouter(lrRouter *gin.RouterGroup) {
	group := lrRouter.Group("lr")
	lrApi := v1.ApiGroupApp.LendingRecordApi
	{
		group.GET("list/:page", lrApi.GetLrList)
		group.GET("list/search/:page", lrApi.GetLrListByCondition)
		group.POST(constant.EMPTY, lrApi.AddLrRec)
		group.PUT(constant.EMPTY, lrApi.UpdateLrRec)
		group.PUT("hs", lrApi.UpdateLrRecHs)
		group.DELETE(constant.EMPTY, lrApi.DelLrRec)
	}
}
