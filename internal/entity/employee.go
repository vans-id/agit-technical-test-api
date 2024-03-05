package entity

import (
	"time"

	"github.com/vans-id/agit-technical-test-api.git/internal/dto"
	"gorm.io/gorm"
)

type Employee struct {
	Id           uint   `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"not null"`
	Nip          string `gorm:"not null"`
	PlaceOfBirth string `gorm:"not null"`
	DateOfBirth  string `gorm:"not null"`
	Age          uint   `gorm:"not null"`
	Address      string `gorm:"not null"`
	Religion     uint   `gorm:"not null"`
	Gender       uint   `gorm:"not null"`
	Phone        string `gorm:"not null"`
	Email        string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

func (e *Employee) ToResponse() *dto.EmployeeResponse {
	return &dto.EmployeeResponse{
		Id:           e.Id,
		Name:         e.Name,
		Nip:          e.Nip,
		PlaceOfBirth: e.PlaceOfBirth,
		DateOfBirth:  e.DateOfBirth,
		Age:          e.Age,
		Address:      e.Address,
		Religion:     e.Religion,
		Gender:       e.Gender,
		Phone:        e.Phone,
		Email:        e.Email,
	}
}
