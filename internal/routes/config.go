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

	// jobRepo := repository.NewJobRepository(db)
	// jobUsecase := usecase.NewJobUsecase(jobRepo)
	// jobHandler := handler.NewJobHandler(jobUsecase)

	config := RouterConfig{
		AuthHandler: authHandler,
	}

	return config
}
