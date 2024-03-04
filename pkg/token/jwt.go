package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/vans-id/agit-technical-test-api.git/shared/constants"
)

type JWTHelper interface {
	GenerateToken(userId uint) (string, error)
}

type JWTHelperImpl struct{}

func NewJWTHelper() JWTHelper {
	return &JWTHelperImpl{}
}

func (h *JWTHelperImpl) GenerateToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":     constants.JWT_ISS,
			"iat":     constants.JWT_IAT,
			"exp":     constants.JWT_EXP,
			"user_id": userId,
		})

	return token.SignedString([]byte(constants.APP_SECRET_KEY))
}
