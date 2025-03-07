package services

import (
	"errors"
	"fmt"
	"log"

	"rest-api/config"
	"rest-api/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(creds models.Credentials) (models.User, error) {
	var existingUser models.User
	result := config.DB.Where("username = ?", creds.Username).First(&existingUser)
	if result.Error == nil {
		return models.User{}, errors.New("пользователь с таким именем уже существует")
	}

	hashedPassword, err := hashPassword(creds.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("ошибка хеширования пароля: %w", err)
	}
	user := models.User{
		Username: creds.Username,
		Password: hashedPassword,
	}

	log.Printf("Создан хеш пароля: %s", hashedPassword)

	result = config.DB.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func DeleteUser(username string) error {
	var user models.User
	result := config.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return errors.New("пользователь не найден")
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("ошибка хеширования пароля: %w", err)
	}
	return string(hash), nil
}
