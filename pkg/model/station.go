package model

import (
	"time"

	"gorm.io/gorm"
)

type Station struct {
	ID         uint64         `gorm:"column:id;primaryKey" json:"id"`
	DivisionId uint64         `gorm:"column:dividion_id" json:"dividion_id"`
	Name       string         `gorm:"column:name;type:varchar(100);not null" json:"name"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}