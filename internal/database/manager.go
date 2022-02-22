package database

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

var (
	db       *gorm.DB
	dbURI    string
	dbDriver string
)

func GetDatabase() *gorm.DB {
	if db == nil {
		err := connect()
		if err != nil {
			log.Error(err)
		} else {
			log.Info("Database connection established")
		}
	}

	return db
}

func SetDriver(driver string) {
	if db != nil {
		log.Warn("Trying to set db driver after db is already initiated, no changes made!")
		return
	}

	dbDriver = driver
}

func SetURI(uri string) {
	if db != nil {
		log.Warn("Trying to set db uri after db is already initiated, no changes made!")
		return
	}

	dbURI = uri
}

func connect() error {
	if dbDriver == "" {
		return errors.New("trying to initiate db without driver being set")
	}

	var err error
	db, err = gorm.Open(sqlite.Open(dbURI), &gorm.Config{})
	return err
}
