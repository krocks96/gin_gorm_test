package entity

import (
	"time"
)

type User struct {
	ID uint64 `gorm:"primary_key"`
	Name string
	CreatedAt time.Time
	UpdatedAt *time.Time
}