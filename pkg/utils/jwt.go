package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTUtil struct {
	Secret []byte
}

func NewJWTUtil(secret string) *JWTUtil {
	return &JWTUtil{Secret: []byte(secret)}
}

func (j *JWTUtil) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.Secret)
}

func (j *JWTUtil) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return j.Secret, nil
	})
	if err != nil {
		return "", errors.New("invalid or expired token")
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims["userID"].(string), nil
}
