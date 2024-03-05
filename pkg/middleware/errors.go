package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleErrors(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			// resp := dto.JSONResponse{
			// 	Message: err.Error(),
			// }

			// switch {
			// case errors.Is(err, apperror.ErrInvalidInput):
			// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			// case errors.Is(err, apperror.ErrInvalidDate):
			// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			// case errors.Is(err, apperror.ErrGreaterDate):
			// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			// case errors.Is(err, apperror.ErrTopupAmount):
			// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			// case errors.Is(err, apperror.ErrTransferAmount):
			// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			// case errors.Is(err, apperror.ErrInvalidMethod):
			// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			// case errors.Is(err, apperror.ErrInvalidOTP):
			// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			// case errors.Is(err, apperror.ErrNoAttemptLeft):
			// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			// case errors.Is(err, apperror.ErrNotFound):
			// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
			// case errors.Is(err, apperror.ErrAlreadyExists):
			// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)

			// case errors.Is(err, apperror.ErrInvalidCred):
			// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			// case errors.Is(err, apperror.ErrNotAuthorized):
			// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			// case errors.Is(err, apperror.ErrMissingMetadata):
			// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			// case errors.Is(err, apperror.ErrWrongUser):
			// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)

			// case errors.Is(err, context.Canceled):
			// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
			// case errors.Is(err, context.DeadlineExceeded):
			// 	ctx.AbortWithStatusJSON(http.StatusGatewayTimeout, resp)

			// case errors.Is(err, apperror.ErrNoDataAvailable):
			// 	ctx.Header("X-Info", err.Error())
			// 	ctx.AbortWithStatusJSON(http.StatusNoContent, resp)

			// default:
			// 	dbResp := dto.JSONResponse{
			// 		Message: apperror.ErrDB.Error(),
			// 	}
			// 	sqlDb, err := db.DB()
			// 	if err != nil {
			// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dbResp)
			// 		return
			// 	}
			// 	if err := sqlDb.Ping(); err != nil {
			// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dbResp)
			// 		return
			// 	}

			// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
			// }
			// }
		}
	}
}
