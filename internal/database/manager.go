package database

import (
	"Claerance/internal/users"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
)

var (
	db       Databaser
	dbURI    string
	dbDriver string
)

type database struct {
	db *sql.DB
}

type Databaser interface {
	createTableFromFile(string)
	AddUser(string, string)
	GetUserByName(string) (users.User, error)
	GetUserById(int) (users.User, error)
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
	d.createTableFromFile("users")

	return d
}

func (d database) createTableFromFile(tablename string) {
	statement, err := ioutil.ReadFile(tablename + ".sql")
	if err != nil {
		log.Printf("ERROR - Could not read file %s", tablename+".sql")
		return
	}

	res, err := d.db.Exec(string(statement))
	if err != nil {
		log.Println("Table already exists, skipping...")
		log.Fatal(err)
	} else {
		log.Println(res.RowsAffected()) // TODO find how to check if success
		log.Printf("Table %s created", tablename)
	}
}

func (d database) AddUser(username string, hashedPW string) {
	user, err := d.GetUserByName(username)

	if err == nil && user.PwdHash != "" {
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

func (d database) GetUserByName(username string) (users.User, error) {
	var u users.User

	query := `SELECT * FROM users WHERE username = ?`
	err := d.db.QueryRow(query, username).Scan(&u.Username, &u.PwdHash, &u.CreatedAt, &u.Email, &u.TelegramId)

	return u, err
}

func (d database) GetUserById(userId int) (users.User, error) {
	var u users.User

	query := `SELECT * FROM users WHERE id = ?`
	err := d.db.QueryRow(query, userId).Scan(&u.Username, &u.PwdHash, &u.CreatedAt, &u.Email, &u.TelegramId)

	return u, err
}
