package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint64 `gorm:"primary_key"`
	DisplayName string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   gorm.DeletedAt
}
