package middlewares

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWTToken(id uint) (string, error) {
	if id == 0 {
		return "", errors.New("Error while generating token")
	}

	key := os.Getenv("JWT_KEY")

	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractJWTToken(token string) (interface{}, error) {
	key := os.Getenv("JWT_KEY")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return 0, err
	}

	return claims["id"], nil
}
