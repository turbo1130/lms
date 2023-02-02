package v1

import (
	"LibraryManagementSys/model"
	"LibraryManagementSys/model/response"
	"LibraryManagementSys/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
}

// Login
// @Tags 系统
// @Summary 登录
// @Description 登录
// @Param user body request.LoginUser true "用户登录名、密码"
// @Router /login [post]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (b *BaseApi) Login(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var user model.User
	_ = json.Unmarshal(data, &user)
	res, err := userService.Login(user)
	if err != nil {
		response.FailWithMessage("登录失败 "+err.Error(), ctx)
	} else {
		response.OkWithData(res, ctx)
	}
}

// Register
// @Tags 系统
// @Summary 账户注册
// @Description 账户注册
// @Param user body model.User true "用户"
// @Router /register [post]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (b *BaseApi) Register(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var user model.User
	_ = json.Unmarshal(data, &user)
	if utils.IsEmpty(user.Password) || utils.IsEmpty(user.Username) {
		response.FailWithMessage("用户名或密码为空，注册失败", ctx)
	} else {
		err := userService.Register(user)
		if err != nil {
			response.FailWithMessage("注册失败 "+err.Error(), ctx)
		} else {
			response.OkWithMessage("注册成功！", ctx)
		}
	}
}
