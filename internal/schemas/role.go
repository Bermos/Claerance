package schemas

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name  string  `json:"name"`
	Users []*User `json:"users" gorm:"many2many:user_roles"`
}

type CreateRoleRequest struct {
	Name string `json:"name"`
}

func CreateRole(request map[string]interface{}) {
	db.Create(&Role{
		Name: request["name"].(string),
	})
}
