package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secretKey []byte
}

type JwtCustomClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func NewJWTManager(secret string) *JWTManager {
	return &JWTManager{
		secretKey: []byte(secret),
	}
}

func (j *JWTManager) GenerateAccessToken(userID uint, role string) (string, error) {

	claims := JwtCustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}
