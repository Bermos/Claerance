package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Site struct {
	gorm.Model
	Name            string    `json:"name"`
	Url             string    `json:"url"`
	FirstContact    time.Time `json:"firstContact"`
	LastContact     time.Time `json:"lastContact"`
	AutoDiscover    bool      `json:"autoDiscover"`
	AuthorizedUsers []*User   `json:"authorizedUsers" gorm:"many2many:site_users"`
	AuthorizedRoles []*Role   `json:"authorizedRoles" gorm:"many2many:site_roles"`
}

type SiteAuthorizedEntities struct {
	Users []*User `json:"users"`
	Roles []*Role `json:"roles"`
}

type CreateSiteRequest struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func CreateSite(request map[string]interface{}) {
	db.Create(&Site{
		Name:         request["name"].(string),
		Url:          request["url"].(string),
		FirstContact: time.Now(),
		LastContact:  time.Now(),
	})
}
