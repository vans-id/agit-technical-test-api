package entity

import (
	"time"

	"github.com/vans-id/agit-technical-test-api.git/internal/dto"
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

func (u *User) ToRegisterResponse() *dto.RegisterResponse {
	return &dto.RegisterResponse{
		Id:       u.Id,
		Username: u.Username,
	}
}
