package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vans-id/agit-technical-test-api.git/internal/dto"
	"github.com/vans-id/agit-technical-test-api.git/shared/apperror"
)

func HandleErrors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			resp := dto.JSONResponse{
				Message: err.Error(),
			}

			switch {
			case errors.Is(err, apperror.ErrInvalidInput):
				ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.Is(err, apperror.ErrNotFound):
				ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			case errors.Is(err, apperror.ErrAlreadyExists):
				ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)

			case errors.Is(err, apperror.ErrInvalidCred):
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			case errors.Is(err, apperror.ErrNotAuthorized):
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			case errors.Is(err, apperror.ErrMissingToken):
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)

			default:
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
			}
		}
	}
}
