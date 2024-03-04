package routes

import (
	"gorm.io/gorm"
)

func GetConfig(db *gorm.DB) RouterConfig {
	// jwtHelper := token.NewJWTHelper()

	// jobRepo := repository.NewJobRepository(db)
	// jobUsecase := usecase.NewJobUsecase(jobRepo)
	// jobHandler := handler.NewJobHandler(jobUsecase)

	// userRepo := repository.NewUserRepository(db)
	// userUsecase := usecase.NewUserUsecase(userRepo, jwtHelper)
	// userHandler := handler.NewUserHandler(userUsecase)

	config := RouterConfig{}

	return config
}
