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
	Url  string `json:"url"`
}

func Setup() {
	db = database.GetDatabase()

	if err := db.AutoMigrate(&Role{}); err != nil {
		log.Println("WARNING - Could not migrate db schema Role")
	}
}

func GetRoleById(roleId int) (Role, error) {
	var role Role
	result := db.First(&role, roleId)
	return role, result.Error
}

func GetRoleByName(name string) (Role, error) {
	var role Role
	result := db.First(&role, "name = ?", name)
	return role, result.Error
}

func GetAllRoles() ([]Role, error) {
	var roleList []Role
	result := db.Find(&roleList)
	return roleList, result.Error
}

func DeleteRole(role Role) bool {
	result := db.Delete(&role)
	return result.RowsAffected == 1
}

func DeleteRoleById(roleId int) bool {
	result := db.Delete(&Role{}, roleId)
	return result.RowsAffected == 1
}

func UpdateRole(role Role) error {
	result := db.Save(&role)
	return result.Error
}
