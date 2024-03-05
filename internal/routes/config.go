package routes

import (
	"github.com/vans-id/agit-technical-test-api.git/internal/handler"
	"github.com/vans-id/agit-technical-test-api.git/internal/repository"
	"github.com/vans-id/agit-technical-test-api.git/internal/usecase"
	"github.com/vans-id/agit-technical-test-api.git/pkg/hasher"
	"github.com/vans-id/agit-technical-test-api.git/pkg/token"
	"gorm.io/gorm"
)

func GetConfig(db *gorm.DB) RouterConfig {
	jwtHelper := token.NewJWTHelper()
	hasher := hasher.NewBcryptHasher()

	authRepo := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo, jwtHelper, hasher)
	authHandler := handler.NewAuthHandler(authUsecase)

	emplRepo := repository.NewEmplRepository(db)
	emplUsecase := usecase.NewEmplUsecase(emplRepo)
	emplHandler := handler.NewEmplHandler(emplUsecase)

	config := RouterConfig{
		AuthHandler:     authHandler,
		EmployeeHandler: emplHandler,
	}

	return config
}
