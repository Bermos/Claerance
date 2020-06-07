package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type database struct {
	url string
	db  *sql.DB
}

type Databaser interface {
	Connect(string)
	AddUser(string, string)
	GetUser(string) (string, error)
}

func NewDatabase(driver string, url string) Databaser {
	db, _ := sql.Open(driver, "test.db")
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connection established")
	}

	_, err := db.Exec(`CREATE TABLE users (id INT AUTO_INCREMENT, username TEXT NOT NULL, password TEXT NOT NULL, PRIMARY KEY (id))`)

	if err != nil {
		if err.Error() == "table users already exists" {
			log.Println("Table already exists, skipping...")
		} else {
			log.Fatal(err)
		}
	} else {
		log.Println("Table users created")
	}

	_, err = db.Exec(`CREATE TABLE sessions (id INT AUTO_INCREMENT, username TEXT NOT NULL, sessionid TEXT NOT NULL, PRIMARY KEY (id))`)

	if err != nil {
		if err.Error() == "table sessions already exists" {
			log.Println("Table already exists, skipping...")
		} else {
			log.Fatal(err)
		}
	} else {
		log.Println("Table sessions created")
	}

	return &database{url, db}
}

func (d database) Connect(url string) {
	d.url = url
	log.Println(d.db == nil)

	d.db, _ = sql.Open("sqlite3", "test.db")
	if err := d.db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("database connection established")
	}

	_, err := d.db.Exec(`CREATE TABLE users (id INT AUTO_INCREMENT, username TEXT NOT NULL, password TEXT NOT NULL, PRIMARY KEY (id))`)

	if err != nil {
		if err.Error() == "table users already exists" {
			log.Println("Table already exists, skipping...")
		} else {
			log.Fatal(err)
		}
	} else {
		log.Println("Table users created")
	}
}

func (d database) AddUser(username string, hashedPW string) {
	pwdHash, err := d.GetUser(username)

	if err == nil && pwdHash != "" {
		log.Println("User not added. Username already exists")
		return
	}

	_, err = d.db.Exec(`INSERT INTO users (username, password) VALUES (?, ?)`, username, hashedPW)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("User added.")
	}
}

func (d database) GetUser(username string) (string, error) {
	var pwdHash string

	query := `SELECT password FROM users WHERE username = ?`
	err := d.db.QueryRow(query, username).Scan(&pwdHash)

	return pwdHash, err
}
