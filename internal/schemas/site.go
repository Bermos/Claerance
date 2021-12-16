package schemas

import (
	"Claerance/internal/database"
	"gorm.io/gorm"
	"time"
)

type Site struct {
	gorm.Model
	Name         string    `json:"name"`
	Url          string    `json:"url"`
	FirstContact time.Time `json:"firstContact"`
	LastContact  time.Time `json:"lastContact"`
}

type CreateSiteRequest struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func CreateSite(request map[string]interface{}) {

	db := database.GetDatabase()
	db.Create(&Site{
		Name:         request["name"].(string),
		Url:          request["url"].(string),
		FirstContact: time.Now(),
		LastContact:  time.Now(),
	})
}
