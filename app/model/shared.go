package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        string         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
