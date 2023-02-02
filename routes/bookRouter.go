package routes

import (
	v1 "LibraryManagementSys/api/v1"
	"LibraryManagementSys/constant"
	"github.com/gin-gonic/gin"
)

type BookRouter struct {
}

func (b *BookRouter) InitBookRouter(bookRouter *gin.RouterGroup) {
	router := bookRouter.Group("/book")
	bookApi := v1.ApiGroupApp.BookApi
	{
		router.GET("list/:page", bookApi.GetBookListByPage)
		router.GET("list/search/:page", bookApi.GetBookListByConditions)
		router.GET("list/publisher", bookApi.GetPublisherList)
		router.POST(constant.EMPTY, bookApi.AddBook)
		router.PUT(constant.EMPTY, bookApi.UpdateBook)
		router.DELETE(constant.EMPTY, bookApi.DelBook)
	}
}
