package database

import (
	"Claerance/internal/users"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
)

var (
	db       Databaser
	dbURI    string
	dbDriver string
)

type tableCountRes struct {
	TableCount int
}

type database struct {
	db *sql.DB
}

type Databaser interface {
	createTableFromFile(string, bool)
	AddUser(string, string) error
	GetUserByName(string) (users.User, error)
	GetUserById(int) (users.User, error)
	GetAllUsers() ([]users.User, error)
	DeleteUserById(userId int) bool
	UpdateUser(users.User) error
	tableExists(string) bool
}

func GetDatabase() Databaser {
	if db == nil {
		db = newDatabase()
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

func newDatabase() Databaser {
	if dbDriver == "" {
		log.Fatal("trying to initiate db without driver being set. Abort.")
		return nil
	}

	db, _ := sql.Open(dbDriver, dbURI)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connection established")
	}

	d := &database{db}
	d.createTableFromFile("users", false)
	d.createTableFromFile("roles", false)
	d.createTableFromFile("sites", false)
	d.createTableFromFile("user_roles", false)
	d.createTableFromFile("user_sites", false)
	d.createTableFromFile("role_sites", false)

	return d
}

func (d database) createTableFromFile(tablename string, overwrite bool) {
	if !overwrite && d.tableExists(tablename) {
		log.Printf("Table '%s' already exists, skipping...", tablename)
		return
	}

	filename := fmt.Sprintf("internal/database/%s/%s.sql", dbDriver, tablename)
	statement, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("ERROR - Could not read file %s", filename)
		return
	}

	_, err = d.db.Exec(string(statement))
	if err != nil {
		log.Printf("Error while creating table '%s'. Aborting...", tablename)
		log.Fatal(err)
	} else {
		log.Printf("Table %s created", tablename)
	}
}

func (d database) tableExists(tablename string) bool {
	query := `SELECT COUNT(name) AS tableCount FROM sqlite_master WHERE type = 'table' AND name = ?`
	row := d.db.QueryRow(query, tablename)
	scanRes := new(tableCountRes)
	err := row.Scan(&scanRes.TableCount)
	return err == nil && scanRes.TableCount == 1
}
