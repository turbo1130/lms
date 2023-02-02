package v1

import (
	"LibraryManagementSys/global"
	"LibraryManagementSys/model"
	"LibraryManagementSys/model/request"
	"LibraryManagementSys/model/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BookApi struct {
}

// GetBookListByPage
// @Tags 图书管理
// @Summary 分页查询图书列表
// @Description 获取图书列表，每次获取11条，分页查询
// @Param page path string true "页码"
// @Param Authorization header string true "Authorization"
// @Router /book/list/{page} [get]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (b *BookApi) GetBookListByPage(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Param("page"))
	res, err := bookService.GetBookListByPage(page)
	if err != nil {
		global.LMS_LOG.Errorf("书集列表查询失败！%s", err)
		response.FailWithMessage("书集列表查询失败！", ctx)
	} else {
		response.OkWithData(res, ctx)
	}
}

// GetBookListByConditions
// @Tags 图书管理
// @Summary 根据条件查询图书列表信息
// @Description 根据条件查询图书列表信息
// @Param Authorization header string true "Authorization"
// @Param page path string true "页码"
// @Param num query string false "书号"
// @Param name query string false "书名"
// @Param author query string false "作者"
// @Param isbn query string false "isbn"
// @Param publisher query string false "出版社"
// @Router /book/list/search/{page} [get]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (b *BookApi) GetBookListByConditions(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Param("page"))
	book := model.Book{
		Num:       ctx.Query("num"),
		Name:      ctx.Query("name"),
		Author:    ctx.Query("author"),
		Isbn:      ctx.Query("isbn"),
		Publisher: ctx.Query("publisher"),
	}
	pageReq := request.PageReq{Page: page, Data: book}
	res, err := bookService.GetBookListByConditions(pageReq)
	if err != nil {
		global.LMS_LOG.Errorf("书集列表查询失败！%s", err)
		response.FailWithMessage("书集列表查询失败！", ctx)
	} else {
		response.OkWithData(res, ctx)
	}
}

// GetPublisherList
// @Tags 图书管理
// @Summary 获取所有出版社名称
// @Description 获取所有出版社名称
// @Param Authorization header string true "Authorization"
// @Router /book/list/publisher [get]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (b *BookApi) GetPublisherList(ctx *gin.Context) {
	res, err := bookService.GetPublisherList()
	if err != nil {
		global.LMS_LOG.Errorf("出版社清单查询失败！%s", err)
		response.FailWithMessage("出版社清单查询失败！", ctx)
	} else {
		response.OkWithData(res, ctx)
	}
}

// AddBook
// @Tags 图书管理
// @Summary 添加图书
// @Description 添加图书
// @Param Authorization header string true "Authorization"
// @Param book body model.Book true "图书"
// @Router /book [post]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (b *BookApi) AddBook(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var book model.Book
	_ = json.Unmarshal(data, &book)
	err := bookService.AddBook(book)
	if err != nil {
		global.LMS_LOG.Errorf("添加图书信息失败！%s", err)
		response.FailWithMessage("添加图书信息失败！", ctx)
	} else {
		response.OkWithMessage("添加图书成功！", ctx)
	}
}

// UpdateBook
// @Tags 图书管理
// @Summary 更新图书
// @Description 更新图书
// @Param Authorization header string true "Authorization"
// @Param book body model.Book true "图书"
// @Router /book [put]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (b *BookApi) UpdateBook(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var book model.Book
	_ = json.Unmarshal(data, &book)
	err := bookService.UpdateBook(book)
	if err != nil {
		global.LMS_LOG.Errorf("更新图书信息失败！%s", err)
		response.FailWithMessage("更新图书信息失败！", ctx)
	} else {
		response.OkWithMessage("更新图书成功！", ctx)
	}
}

// DelBook
// @Tags 图书管理
// @Summary 删除图书
// @Description 删除图书
// @Param Authorization header string true "Authorization"
// @Param book body model.Book true "图书"
// @Router /book [delete]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (b *BookApi) DelBook(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var book model.Book
	_ = json.Unmarshal(data, &book)
	err := bookService.DelBook(book)
	if err != nil {
		global.LMS_LOG.Errorf("删除图书信息失败！%s", err)
		response.FailWithMessage("删除图书信息失败！", ctx)
	} else {
		response.OkWithMessage("删除图书成功！", ctx)
	}
}
