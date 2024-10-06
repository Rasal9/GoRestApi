package services

import (
	"errors"

	"rest-api/config"
	"rest-api/models"
)

func CreateUser(creds models.Credentials) (models.User, error) {
	var existingUser models.User
	result := config.DB.Where("username = ?", creds.Username).First(&existingUser)
	if result.Error == nil {
		return models.User{}, errors.New("пользователь с таким именем уже существует")
	}

	hashedPassword := hashPassword(creds.Password)
	user := models.User{
		Username: creds.Username,
		Password: hashedPassword,
	}

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

func hashPassword(password string) string {
	return password
}
