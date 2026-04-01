package config

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/imperionite/reeisto/backend/models"
)

func BootstrapAdmin(db *gorm.DB) {

	username := App.BootstrapAdminUser
	password := App.BootstrapAdminPassword

	if username == "" || password == "" {
		return
	}

	var existing models.User
	if err := db.Where("username = ?", username).First(&existing).Error; err == nil {
		log.Println("Admin already exists")
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	admin := models.User{
		Username: username,
		Password: string(hashed),
		Role:     "admin",
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Println("Failed to create admin:", err)
		return
	}

	log.Println("Admin bootstrapped")
}
