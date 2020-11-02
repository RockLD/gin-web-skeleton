package model

import (
	"time"
)

//数据库基类
type BaseModel struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"_"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"_"`
}
