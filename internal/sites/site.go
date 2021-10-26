package sites

import (
	"Claerance/internal/database"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	db *gorm.DB
)

type Site struct {
	gorm.Model
	Name         string    `json:"name"`
	Url          string    `json:"url"`
	FirstContact time.Time `json:"first_contact"`
	LastContact  time.Time `json:"last_contact"`
}

func Setup() {
	db = database.GetDatabase()

	if err := db.AutoMigrate(&Site{}); err != nil {
		log.Println("WARNING - Could not migrate db schema Site")
	}
}

func CreateSite(name string, url string) {
	db.Create(&Site{
		Name:         name,
		Url:          url,
		FirstContact: time.Now(),
		LastContact:  time.Now(),
	})
}

func GetSiteById(siteId int) (Site, error) {
	var site Site
	result := db.First(&site, siteId)
	return site, result.Error
}

func GetSiteByName(name string) (Site, error) {
	var site Site
	result := db.First(&site, "name = ?", name)
	return site, result.Error
}

func GetAllSites() ([]Site, error) {
	var siteList []Site
	result := db.Find(&siteList)
	return siteList, result.Error
}

func DeleteSite(site Site) bool {
	result := db.Delete(&site)
	return result.RowsAffected == 1
}

func DeleteSiteById(siteId int) bool {
	result := db.Delete(&Site{}, siteId)
	return result.RowsAffected == 1
}

func UpdateSite(site Site) error {
	result := db.Save(&site)
	return result.Error
}
