package config

import (
	"go-auth/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// global DB connection holder
var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	DB = db
}
