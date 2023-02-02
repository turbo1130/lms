package model

import "LibraryManagementSys/global"

type Role struct {
	global.LmsModel
	RoleEn   string `json:"roleEn" gorm:"column:roleEn;comment:角色代码"`
	RoleName string `json:"roleName" gorm:"column:roleName;comment:角色名"`
}
