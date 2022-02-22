package schemas

import (
	"Claerance/internal/database"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

func Setup() {
	db = database.GetDatabase()

	if err := db.AutoMigrate(&Role{}); err != nil {
		log.Warn("Could not migrate db schema Role")
	}

	if err := db.AutoMigrate(&Group{}); err != nil {
		log.Warn("Could not migrate db schema Group")
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		log.Warn("Could not migrate db schema User")
	}

	if err := db.AutoMigrate(&Site{}); err != nil {
		log.Warn("Could not migrate db schema Site")
	}

	// TODO: replace with init user creation
	CreateUser(map[string]interface{}{
		"username": "Admin",
		"password": "admin",
	})
}
