package v1

import (
	"LibraryManagementSys/constant"
	"LibraryManagementSys/model/response"
	"github.com/gin-gonic/gin"
)

type JwtInvalidApi struct {
}

// Logout
// @Tags 系统
// @Summary 退出登录
// @Description 退出登录
// @Param Authorization header string true "Authorization"
// @Router /logout [post]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (b *JwtInvalidApi) Logout(ctx *gin.Context) {
	token := ctx.GetHeader(constant.HeaderAuthorization)
	token = token[7:]
	err := userService.Logout(token)
	if err != nil {
		response.FailWithMessage("退出登录失败 "+err.Error(), ctx)
	} else {
		response.OkWithMessage("退出登录成功！", ctx)
	}
}
