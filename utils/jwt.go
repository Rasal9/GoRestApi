package utils

import (
	"rest-api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("HS256")

func GenerateJWT(username string) (string, time.Time, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &models.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, expirationTime, err
}
