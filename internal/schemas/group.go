package schemas

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name  string  `json:"name"`
	Users []*User `json:"users" gorm:"many2many:user_groups"`
}

type CreateGroupRequest struct {
	Name string `json:"name"`
}

func CreateGroup(request map[string]interface{}) {
	db.Create(&Group{
		Name: request["name"].(string),
	})
}
