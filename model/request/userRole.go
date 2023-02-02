package request

type UserRole struct {
	Id     string `json:"id"`                          // 用户ID
	RoleEn string `gorm:"column:roleEn" json:"roleEn"` // 角色标示
	RoleId string `gorm:"column:roleId" json:"roleId"` // 角色ID
}
