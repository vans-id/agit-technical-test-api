package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vans-id/agit-technical-test-api.git/internal/dto"
	"github.com/vans-id/agit-technical-test-api.git/internal/usecase"
	"github.com/vans-id/agit-technical-test-api.git/pkg/apperror"
)

type AuthHandler interface {
	HandleRegister(ctx *gin.Context)
	HandleLogin(ctx *gin.Context)
}

type authHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(uc usecase.AuthUsecase) AuthHandler {
	return &authHandler{authUsecase: uc}
}

func (h *authHandler) HandleRegister(ctx *gin.Context) {
	param := dto.RegisterRequest{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.Error(apperror.ErrInvalidInput)
		return
	}
}

func (h *authHandler) HandleLogin(ctx *gin.Context) {}
