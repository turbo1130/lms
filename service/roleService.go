package service

import (
	"LibraryManagementSys/global"
	"LibraryManagementSys/model"
)

type RoleService struct {
}

func (r *RoleService) GetRoleList() ([]model.Role, error) {
	var roleList []model.Role
	err := global.LMS_DB.Model(&model.Role{}).Where("deletedAt is NULL").Find(&roleList).Error
	if err != nil {
		return nil, err
	}
	return roleList, nil
}
