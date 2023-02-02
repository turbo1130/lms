package v1

import (
	"LibraryManagementSys/constant"
	"LibraryManagementSys/global"
	"LibraryManagementSys/model"
	"LibraryManagementSys/model/request"
	"LibraryManagementSys/model/response"
	"LibraryManagementSys/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

type LendingRecordApi struct {
}

// GetLrList
// @Tags 借还管理
// @Summary 获取借阅信息列表
// @Description 获取借阅信息列表
// @Param Authorization header string true "Authorization"
// @Param page path string true "页码"
// @Router /lr/list/{page} [get]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (lr *LendingRecordApi) GetLrList(ctx *gin.Context) {
	pageStr := ctx.Param("page")
	page, _ := strconv.Atoi(pageStr)
	res, err := lendingRecordService.GetLrListByPage(page)
	if err != nil {
		response.FailWithMessage("获取借阅信息列表失败！", ctx)
		global.LMS_LOG.Errorf("获取借阅信息列表失败！Err: %s", err.Error())
	} else {
		response.OkWithData(res, ctx)
	}
}

// GetLrListByCondition
// @Tags 借还管理
// @Summary 根据条件查询
// @Description 根据条件查询
// @Param Authorization header string true "Authorization"
// @Param xh query string false "学号"
// @Param jyr query string false "借阅人"
// @Param jysj query string false "借阅时间 yyyy-MM-dd HH:mm:ss"
// @Param bookName query string false "书名"
// @Param isbn query string false "ISBN"
// @Param page path string true "页码"
// @Router /lr/list/search/{page} [get]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (lr *LendingRecordApi) GetLrListByCondition(ctx *gin.Context) {
	param := ctx.Param("page")
	page, _ := strconv.Atoi(param)
	lrCondition := request.LendingRecCondition{
		Isbn:     ctx.Query("isbn"),
		Xh:       ctx.Query("xh"),
		Jyr:      ctx.Query("jyr"),
		BookName: ctx.Query("bookName"),
		Jysj:     ctx.Query("jysj"),
		Page:     page,
	}
	res, err := lendingRecordService.GetLrListByCondition(lrCondition)
	if err != nil {
		response.FailWithMessage("获取借阅信息列表失败！", ctx)
		global.LMS_LOG.Errorf("获取借阅信息列表失败！Err: %s", err.Error())
	} else {
		response.OkWithData(res, ctx)
	}
}

// AddLrRec
// @Tags 借还管理
// @Summary 添加借阅记录
// @Description 添加借阅记录
// @Param Authorization header string true "Authorization"
// @Param borrowBook body request.BorrowBook true "借阅信息"
// @Router /lr [post]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (lr *LendingRecordApi) AddLrRec(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var lrRec model.LendingRecord
	_ = json.Unmarshal(data, &lrRec)
	lrRec.ID = utils.GetRandomId()
	lrRec.IsH = false
	claim := ctx.MustGet(constant.CustomClaimsKey).(*request.CustomClaims)
	lrRec.JyzhId = claim.BaseClaims.ID
	if utils.IsEmpty(lrRec.Jydjr) {
		lrRec.Jydjr = claim.BaseClaims.Username
	}
	err := lendingRecordService.AddLrRec(lrRec)
	if err != nil {
		global.LMS_LOG.Errorf("添加借阅信息失败！%s", err)
		response.FailWithMessage("添加借阅信息失败！", ctx)
	} else {
		response.OkWithMessage("添加借阅成功！", ctx)
	}
}

// UpdateLrRec
// @Tags 借还管理
// @Summary 更新借阅记录
// @Description 更新借阅记录
// @Param Authorization header string true "Authorization"
// @Param lendingRecord body request.BorrowBook true "借阅信息"
// @Router /lr [put]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (lr *LendingRecordApi) UpdateLrRec(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var lrRec model.LendingRecord
	_ = json.Unmarshal(data, &lrRec)
	lrRec.BookId = constant.EMPTY
	err := lendingRecordService.UpdateLrRec(lrRec)
	if err != nil {
		global.LMS_LOG.Errorf("更新借阅信息失败！%s", err)
		response.FailWithMessage("更新借阅信息失败！", ctx)
	} else {
		response.OkWithMessage("更新借阅成功！", ctx)
	}
}

// UpdateLrRecHs
// @Tags 借还管理
// @Summary 还书
// @Description 还书
// @Param Authorization header string true "Authorization"
// @Param returnBook body request.ReturnBook true "还书信息"
// @Router /lr/hs [put]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (lr *LendingRecordApi) UpdateLrRecHs(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var lrRec model.LendingRecord
	_ = json.Unmarshal(data, &lrRec)
	lrRec.IsH = true
	claim := ctx.MustGet(constant.CustomClaimsKey).(*request.CustomClaims)
	lrRec.HszhId = claim.BaseClaims.ID
	if utils.IsEmpty(lrRec.Hsdjr) {
		lrRec.Hsdjr = claim.BaseClaims.Username
	}
	err := lendingRecordService.UpdateLrRecHs(lrRec)
	if err != nil {
		global.LMS_LOG.Errorf("还书失败！%s", err)
		response.FailWithMessage("还书失败！", ctx)
	} else {
		response.OkWithMessage("还书成功！", ctx)
	}
}

// DelLrRec
// @Tags 借还管理
// @Summary 删除借阅记录
// @Description 删除借阅记录
// @Param Authorization header string true "Authorization"
// @Param borrowBook body request.BorrowBook true "借阅记录ID"
// @Router /lr [delete]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (lr *LendingRecordApi) DelLrRec(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var lrRec model.LendingRecord
	_ = json.Unmarshal(data, &lrRec)
	err := lendingRecordService.DelLrRec(lrRec.ID)
	if err != nil {
		global.LMS_LOG.Errorf("还书失败！%s", err)
		response.FailWithMessage("删除借阅记录失败！", ctx)
	} else {
		response.OkWithMessage("删除借阅记录成功！", ctx)
	}
}
