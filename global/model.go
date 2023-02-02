package global

import (
	"gorm.io/gorm"
	"time"
)

type LmsModel struct {
	ID        string         `gorm:"column:id;not null;unique;primary_key;comment:ID;size:32" json:"id"` // ID
	CreatedAt time.Time      `gorm:"column:createdAt;comment:创建时间" json:"createdAt"`                     // 创建时间
	UpdatedAt time.Time      `gorm:"column:updatedAt;comment:更新时间" json:"updatedAt"`                     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt;index;comment:删除标记" json:"deletedAt"`               // 删除时间
}
