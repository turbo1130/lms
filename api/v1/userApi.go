package v1

import (
	"LibraryManagementSys/model/request"
	"LibraryManagementSys/model/response"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

// GetUserInfo
// @Tags 用户管理
// @Summary 查询用户信息
// @Description 查询用户信息
// @Param Authorization header string true "Authorization"
// @Param id query string true "用户ID"
// @Router /user/self [get]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (u *UserApi) GetUserInfo(ctx *gin.Context) {
	id := ctx.Query("id")
	userInfo, err := userService.GetUserInfo(id)
	if err != nil {
		response.FailWithMessage("查询用户信息失败 "+err.Error(), ctx)
	} else {
		response.OkWithData(userInfo, ctx)
	}
}

// UpdateUserInfo
// @Tags 用户管理
// @Summary 更新用户信息
// @Description 更新用户信息
// @Param Authorization header string true "Authorization"
// @Param userInfo body request.UserInfo true "用户信息：password、changeTime字段无效"
// @Router /user/self [put]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (u *UserApi) UpdateUserInfo(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var userInfo request.UserInfo
	_ = json.Unmarshal(data, &userInfo)
	err := userService.UpdateUserInfo(userInfo)
	if err != nil {
		response.FailWithMessage("更新用户信息失败 "+err.Error(), ctx)
	} else {
		response.OkWithMessage("更新用户信息成功", ctx)
	}
}

// ChangeUserPassword
// @Tags 用户管理
// @Summary 更改密码
// @Description 更改密码
// @Param Authorization header string true "Authorization"
// @Param userInfo body request.UserInfo true "用户信息：只有id、password、changeTime字段有效"
// @Router /user/self [post]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (u *UserApi) ChangeUserPassword(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var userInfo request.UserInfo
	_ = json.Unmarshal(data, &userInfo)
	err := userService.ChangeUserPassword(userInfo.Id, userInfo.Password, userInfo.ChangeTime)
	if err != nil {
		response.FailWithMessage("更改密码失败 "+err.Error(), ctx)
	} else {
		response.OkWithMessage("更改密码成功！", ctx)
	}
}
