package services

import (
	"errors"
	"rest-api/config"
	"rest-api/models"

	"golang.org/x/crypto/bcrypt"
)

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AuthenticateUser(creds models.Credentials) (models.User, error) {
	var user models.User
	result := config.DB.Where("username = ?", creds.Username).First(&user)

	if result.Error != nil {
		return models.User{}, errors.New("неверные учетные данные")
	}

	if !checkPasswordHash(creds.Password, user.Password) {
		return models.User{}, errors.New("неверные учетные данные")
	}

	return user, nil
}
