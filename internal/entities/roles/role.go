package roles

import (
	"Claerance/internal/database"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

type Role struct {
	gorm.Model
	Name string `json:"name"`
}

type CreateRoleRequest struct {
	Name string `json:"name"`
}

func Setup() {
	db = database.GetDatabase()

	if err := db.AutoMigrate(&Role{}); err != nil {
		log.Println("WARNING - Could not migrate db schema Role")
	}
}

func CreateRole(request map[string]interface{}) {
	db := database.GetDatabase()
	db.Create(&Role{
		Name: request["name"].(string),
	})
}
