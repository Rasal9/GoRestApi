package models

import "github.com/dgrijalva/jwt-go"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
