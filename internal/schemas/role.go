package schemas

import (
	"Claerance/internal/database"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string `json:"name"`
}

type CreateRoleRequest struct {
	Name string `json:"name"`
}

func CreateRole(request map[string]interface{}) {
	db := database.GetDatabase()
	db.Create(&Role{
		Name: request["name"].(string),
	})
}
