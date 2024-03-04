package migration

import (
	"github.com/vans-id/agit-technical-test-api.git/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	users := []*entity.User{
		{
			Username: "admin",
			Password: hashPassword("admin"),
		},
	}
	db.Create(users)
}

func hashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword)
}
