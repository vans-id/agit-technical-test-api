package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vans-id/agit-technical-test-api.git/shared/apperror"
	"github.com/vans-id/agit-technical-test-api.git/shared/constants"
)

func HandleAuthentication(ctx *gin.Context) {
	var secretKey = constants.APP_SECRET_KEY

	tokenString := ctx.Request.Header.Get("Authorization")
	if tokenString == "" {
		ctx.Error(apperror.ErrMissingToken)
		ctx.Abort()
		return
	}
	tokenString = tokenString[len("Bearer "):]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		ctx.Error(apperror.ErrNotAuthorized)
		ctx.Abort()
		return
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	ctx.Set(constants.UserId, userId)
	ctx.Next()
}
