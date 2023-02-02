package v1

import (
	"LibraryManagementSys/model/request"
	"LibraryManagementSys/model/response"
	"LibraryManagementSys/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type RoleApi struct {
}

// UpdateUserRole
// @Tags 角色
// @Summary 更改用户角色
// @Description 更改用户角色
// @Param Authorization header string true "Authorization"
// @Param role body request.UserRole true "角色信息"
// @Router /role/ur [put]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (r *RoleApi) UpdateUserRole(ctx *gin.Context) {
	data, _ := ctx.GetRawData()
	var roleUr request.UserRole
	_ = json.Unmarshal(data, &roleUr)
	if utils.IsEmpty(roleUr.Id) || utils.IsEmpty(roleUr.RoleEn) || utils.IsEmpty(roleUr.RoleId) {
		response.FailWithMessage("更新用户角色权限失败: 信息不全", ctx)
		return
	}
	err := userService.UpdateUserRole(roleUr)
	if err != nil {
		response.FailWithMessage("更新用户角色权限失败 "+err.Error(), ctx)
	} else {
		response.OkWithMessage("更新用户角色权限成功", ctx)
	}
}

// GetRoleList
// @Tags 角色
// @Summary 获取角色信息列表
// @Description 获取角色信息列表
// @Param Authorization header string true "Authorization"
// @Router /role/list [get]
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
func (r *RoleApi) GetRoleList(ctx *gin.Context) {
	roleList, err := roleService.GetRoleList()
	if err != nil {
		response.FailWithMessage("获取角色信息列表失败 "+err.Error(), ctx)
	} else {
		response.OkWithData(roleList, ctx)
	}
}
