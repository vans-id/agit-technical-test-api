package repository

import (
	"context"
	"errors"

	"github.com/vans-id/agit-technical-test-api.git/internal/entity"
	"gorm.io/gorm"
)

type EmplRepository interface {
	Find(ctx context.Context, search string) ([]*entity.Employee, error)
	FindOneById(ctx context.Context, id uint) (*entity.Employee, error)
	Create(ctx context.Context, payload *entity.Employee) error
	Update(ctx context.Context, payload *entity.Employee) error
	Delete(ctx context.Context, payload *entity.Employee) error
}

type emplRepository struct {
	db *gorm.DB
}

func NewEmplRepository(db *gorm.DB) EmplRepository {
	return &emplRepository{db: db}
}

func (r *emplRepository) Find(ctx context.Context, search string) ([]*entity.Employee, error) {
	employees := []*entity.Employee{}

	err := r.db.Where(`name ILIKE ?`, "%"+search+"%").Find(&employees).Error
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *emplRepository) FindOneById(ctx context.Context, id uint) (*entity.Employee, error) {
	employees := entity.Employee{}

	err := r.db.Where(`id = ?`, id).First(&employees).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, nil
		default:
			return nil, err
		}
	}

	return &employees, nil
}

func (r *emplRepository) Create(ctx context.Context, newEmployee *entity.Employee) error {
	return r.db.Create(&newEmployee).Error
}

func (r *emplRepository) Update(ctx context.Context, employee *entity.Employee) error {
	return r.db.Save(&employee).Error
}

func (r *emplRepository) Delete(ctx context.Context, employee *entity.Employee) error {
	return r.db.Delete(&employee).Error
}
