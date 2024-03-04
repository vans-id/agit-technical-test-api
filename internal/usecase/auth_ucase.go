package usecase

import (
	"context"

	"github.com/vans-id/agit-technical-test-api.git/internal/entity"
	"github.com/vans-id/agit-technical-test-api.git/internal/repository"
	"github.com/vans-id/agit-technical-test-api.git/pkg/hasher"
	"github.com/vans-id/agit-technical-test-api.git/pkg/token"
	"github.com/vans-id/agit-technical-test-api.git/shared/apperror"
)

type AuthUsecase interface {
	Register(ctx context.Context, payload *entity.User) (*entity.User, error)
	Login(ctx context.Context, payload *entity.User) (string, error)
}

type authUsecase struct {
	authRepo repository.AuthRepository
	jwt      token.JWTHelper
	hasher   hasher.Hasher
}

func NewAuthUsecase(r repository.AuthRepository, jwt token.JWTHelper, hasher hasher.Hasher) AuthUsecase {
	return &authUsecase{authRepo: r, jwt: jwt, hasher: hasher}
}

func (uc *authUsecase) Register(ctx context.Context, payload *entity.User) (*entity.User, error) {
	payload, err := uc.authRepo.FindOneByUsername(ctx, payload.Username)
	if err != nil {
		return nil, err
	}
	if payload != nil {
		return nil, apperror.ErrAlreadyExists
	}

	hashedPassword, err := uc.hasher.HashPassword([]byte(payload.Password))
	if err != nil {
		return nil, err
	}
	payload.Password = string(hashedPassword)

	err = uc.authRepo.Create(ctx, payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (uc *authUsecase) Login(ctx context.Context, payload *entity.User) (string, error) {
	user, err := uc.authRepo.FindOneByUsername(ctx, payload.Username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", apperror.ErrInvalidCred
	}

	err = uc.hasher.ComparePasswords([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return "", apperror.ErrInvalidCred
	}

	tokenString, err := uc.jwt.GenerateToken(user.Id)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
