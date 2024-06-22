package models

import (
	"gorm.io/gorm"
	"time"
)

// ID 自增ID主键
type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

// Timestamps 创建、更新时间
type Timestamps struct {
	CreatedBy time.Time `json:"created_by"`
	UpdatedBy time.Time `json:"updated_by"`
}

// SoftDeletes 软删除
type SoftDeletes struct {
	Deleted gorm.DeletedAt `json:"deleted" gorm:"index"`
}
