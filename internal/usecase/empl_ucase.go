package usecase

import (
	"context"

	"github.com/vans-id/agit-technical-test-api.git/internal/entity"
	"github.com/vans-id/agit-technical-test-api.git/internal/repository"
	"github.com/vans-id/agit-technical-test-api.git/pkg/apperror"
)

type EmplUsecase interface {
	GetAll(ctx context.Context, search string) ([]*entity.Employee, error)
	GetDetail(ctx context.Context, id uint) (*entity.Employee, error)
	Add(ctx context.Context, payload *entity.Employee) error
	Edit(ctx context.Context, payload *entity.Employee) error
	Remove(ctx context.Context, payload *entity.Employee) error
}

type emplUsecase struct {
	emplRepo repository.EmplRepository
}

func NewEmplUsecase(r repository.EmplRepository) EmplUsecase {
	return &emplUsecase{emplRepo: r}
}

func (uc *emplUsecase) GetAll(ctx context.Context, search string) ([]*entity.Employee, error) {
	return uc.emplRepo.Find(ctx, search)
}

func (uc *emplUsecase) GetDetail(ctx context.Context, id uint) (*entity.Employee, error) {
	return uc.emplRepo.FindOneById(ctx, id)
}

func (uc *emplUsecase) Add(ctx context.Context, payload *entity.Employee) error {
	return uc.emplRepo.Create(ctx, payload)
}

func (uc *emplUsecase) Edit(ctx context.Context, payload *entity.Employee) error {
	currentEmployee, err := uc.emplRepo.FindOneById(ctx, payload.Id)
	if err != nil {
		return err
	}
	if currentEmployee == nil {
		return apperror.ErrInvalidInput
	}

	currentEmployee.Name = payload.Name
	currentEmployee.Nip = payload.Nip
	currentEmployee.PlaceOfBirth = payload.PlaceOfBirth
	currentEmployee.DateOfBirth = payload.DateOfBirth
	currentEmployee.Age = payload.Age
	currentEmployee.Address = payload.Address
	currentEmployee.Religion = payload.Religion
	currentEmployee.Gender = payload.Gender
	currentEmployee.Phone = payload.Phone
	currentEmployee.Email = payload.Email

	return uc.emplRepo.Update(ctx, currentEmployee)
}

func (uc *emplUsecase) Remove(ctx context.Context, payload *entity.Employee) error {
	currentEmployee, err := uc.emplRepo.FindOneById(ctx, payload.Id)
	if err != nil {
		return err
	}
	if currentEmployee == nil {
		return apperror.ErrInvalidInput
	}

	return uc.emplRepo.Delete(ctx, currentEmployee)
}
