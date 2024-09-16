package models

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func GenerateJWT(id uint) (string, error) {
	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &Claims{
		Id: string(id),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func DecodeToken(tokenString string) (*Claims, error) {
	// Инициализируем переменную для claims
	claims := &Claims{}

	// Парсим токен
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Убедитесь, что метод подписи токена - HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Возвращаем секретный ключ для проверки подписи токена
		return &jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Проверяем, валиден ли токен
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
