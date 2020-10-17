package database

import (
	"Claerance/internal/users"
	"database/sql"
	"errors"
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

type database struct {
	db *sql.DB
}

type Databaser interface {
	createTableFromFile(string)
	AddUser(string, string) error
	GetUserByName(string) (users.User, error)
	GetUserById(int) (users.User, error)
	GetAllUsers() ([]users.User, error)
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
	filename := fmt.Sprintf("internal/database/%s/%s.sql", dbDriver, tablename)
	statement, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("ERROR - Could not read file %s", tablename+".sql")
		return
	}

	res, err := d.db.Exec(string(statement))
	if err != nil {
		log.Println("Table already exists, skipping...")
		log.Fatal(err)
	} else {
		log.Println(res.LastInsertId()) // TODO find how to check if success
		log.Printf("Table %s created", tablename)
	}
}

func (d database) AddUser(username string, hashedPW string) error {
	user, err := d.GetUserByName(username)

	if err == nil && user.PwdHash != "" {
		return errors.New("username already exists")
	}

	_, err = d.db.Exec(`INSERT INTO users (username, password) VALUES (?, ?)`, username, hashedPW)
	return err
}

func (d database) GetUserByName(username string) (users.User, error) {
	query := `SELECT * FROM users WHERE username = ?`
	row := d.db.QueryRow(query, username)

	return getUserFromQueryRow(row)
}

func (d database) GetUserById(userId int) (users.User, error) {
	query := `SELECT * FROM users WHERE id = ?`
	row := d.db.QueryRow(query, userId)

	return getUserFromQueryRow(row)
}

func (d database) GetAllUsers() ([]users.User, error) {
	query := `SELECT * FROM users`
	rows, err := d.db.Query(query)

	var us []users.User
	for rows.Next() {
		user, _ := getUserFromQueryRows(rows)
		us = append(us, user)
	}

	return us, err
}

func getUserFromQueryRow(row *sql.Row) (users.User, error) {
	var u users.User
	var email sql.NullString
	var telegramID sql.NullInt32

	err := row.Scan(&u.Id, &u.Username, &u.PwdHash, &u.CreatedAt, &email, &telegramID)

	if email.Valid {
		u.Email = email.String
	}

	if telegramID.Valid {
		u.TelegramId = int(telegramID.Int32)
	}

	return u, err
}

func getUserFromQueryRows(row *sql.Rows) (users.User, error) {
	var u users.User
	var email sql.NullString
	var telegramID sql.NullInt32

	err := row.Scan(&u.Id, &u.Username, &u.PwdHash, &u.CreatedAt, &email, &telegramID)

	if email.Valid {
		u.Email = email.String
	}

	if telegramID.Valid {
		u.TelegramId = int(telegramID.Int32)
	}

	return u, err
}
