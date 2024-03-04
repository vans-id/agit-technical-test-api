package token

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vans-id/agit-technical-test-api.git/pkg/constants"
)

type JWTHelper interface {
	GenerateToken(userId uint) (string, error)
}

type JWTHelperImpl struct{}

func NewJWTHelper() JWTHelper {
	return &JWTHelperImpl{}
}

func (h *JWTHelperImpl) GenerateToken(userId uint) (string, error) {
	var secretKey = os.Getenv("APP_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":     constants.JWT_ISS,
			"iat":     constants.JWT_IAT,
			"exp":     constants.JWT_EXP,
			"user_id": userId,
		})

	return token.SignedString([]byte(secretKey))
}
