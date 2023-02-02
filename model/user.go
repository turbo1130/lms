package model

import (
	"LibraryManagementSys/global"
	"time"
)

type User struct {
	global.LmsModel
	Username  string    `gorm:"column:username;unique;comment:用户名" json:"username"`
	Password  string    `gorm:"column:password;comment:密码" json:"password"`
	RoleEn    string    `gorm:"column:roleEn;size:50;comment:角色" json:"roleEn"`
	RoleId    string    `gorm:"column:roleId;comment:角色ID" json:"roleId"`
	Phone     string    `gorm:"column:phone;comment:用户手机号" json:"phone"`
	Email     string    `gorm:"column:email;comment:用户邮箱" json:"email"`
	NickName  string    `gorm:"column:nickName;default:系统用户;comment:用户昵称" json:"nickName"`
	LoginTime time.Time `gorm:"-" json:"loginTime"`
}
