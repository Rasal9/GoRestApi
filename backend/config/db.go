package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func InitDB() {
	var err error

	dsn := "host=db user=user password=password dbname=authdb port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to initialize database, got error %v", err)
	}

	err = DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("failed to migrate database, got error %v", err)
	}

	log.Println("Database connection and migration successful")
}
