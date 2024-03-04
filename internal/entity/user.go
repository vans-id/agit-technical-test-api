package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
