package schemas

import (
	"Claerance/internal/database"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func Setup() {
	db = database.GetDatabase()

	if err := db.AutoMigrate(&Role{}); err != nil {
		log.Println("WARNING - Could not migrate db schema Role")
	}

	if err := db.AutoMigrate(&Site{}); err != nil {
		log.Println("WARNING - Could not migrate db schema Site")
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		log.Println("WARNING - Could not migrate db schema User")
	}

	// TODO: replace with init user creation
	CreateUser(map[string]interface{}{
		"username": "Admin",
		"password": "admin",
	})
}
