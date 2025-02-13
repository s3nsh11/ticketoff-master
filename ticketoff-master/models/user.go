package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&User{})
}
