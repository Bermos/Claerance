package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var (
	db       *gorm.DB
	dbURI    string
	dbDriver string
)

func GetDatabase() *gorm.DB {
	if db == nil {
		Connect()
	}

	return db
}

func SetDriver(driver string) {
	if db != nil {
		log.Println("WARNING - trying to set db driver after db is already initiated, no changes made!")
		return
	}

	dbDriver = driver
}

func SetURI(uri string) {
	if db != nil {
		log.Println("WARNING - trying to set db uri after db is already initiated, no changes made!")
		return
	}

	dbURI = uri
}

func Connect() {
	if dbDriver == "" {
		log.Fatal("ERROR - trying to initiate db without driver being set. Abort.")
	}

	var err error
	db, err = gorm.Open(sqlite.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("INFO - Database connection established")
	}
}
