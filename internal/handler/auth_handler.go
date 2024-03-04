package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vans-id/agit-technical-test-api.git/internal/dto"
	"github.com/vans-id/agit-technical-test-api.git/internal/entity"
	"github.com/vans-id/agit-technical-test-api.git/internal/usecase"
	"github.com/vans-id/agit-technical-test-api.git/shared/apperror"
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
	param := dto.AuthRequest{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil || strings.Trim(param.Password, " ") == "" {
		ctx.Error(apperror.ErrInvalidInput)
		return
	}
	createdUser, err := h.authUsecase.Register(ctx.Request.Context(), &entity.User{
		Username: param.Username,
		Password: param.Password,
	})
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.JSONResponse{
		Data: createdUser.ToRegisterResponse(),
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (h *authHandler) HandleLogin(ctx *gin.Context) {
	param := dto.AuthRequest{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.Error(apperror.ErrInvalidInput)
		return
	}
	accessToken, err := h.authUsecase.Login(ctx.Request.Context(), &entity.User{
		Username: param.Username,
		Password: param.Password,
	})
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.JSONResponse{
		Data: gin.H{"token": accessToken},
	}
	ctx.JSON(http.StatusOK, resp)
}
