package repository

import (
	"context"
	"errors"

	"github.com/vans-id/agit-technical-test-api.git/internal/entity"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindOneByUsername(ctx context.Context, email string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) FindOneByUsername(ctx context.Context, username string) (*entity.User, error) {
	user := entity.User{}
	err := r.db.
		Where("username = ?", username).First(&user).Error

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, nil
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (r *authRepository) Create(ctx context.Context, user *entity.User) error {
	return r.db.Create(user).Error
}
