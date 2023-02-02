package model

import "LibraryManagementSys/global"

type Version struct {
	global.LmsModel
	Version string `gorm:"column:version;unique;comment:版本号" json:"version"`
}

func (Version) TableName() string {
	return "version"
}
