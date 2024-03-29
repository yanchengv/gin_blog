package global

import (
	"time"

	"gorm.io/gorm"
)

type INIT_COLUMNS struct {
	ID        uint           `gorm:"primarykey" json:"id"` //主键ID
	CreatedAt time.Time      `json:"created_at"`           //创建时间
	UpdatedAt time.Time      `json:"updated_at"`           //更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       //删除时间
}
