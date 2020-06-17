package users

import (
	"Claerance/internal/database"
	"golang.org/x/crypto/bcrypt"
)

var (
	db database.Databaser
)

func Setup() {
	db = database.GetDatabase()

	CreateUser("bermos", "test")
	CreateUser("test", "test")
}

func CreateUser(username string, password string) {
	hashedPW, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	db.AddUser(username, string(hashedPW))
}

func GetUserById(userId int) (User, error) {
	var u User
	var err error
	u, err = db.GetUserById(userId)

	return u, err
}

func GetUserByName(username string) (User, error) {
	var u User
	var err error
	u.PwdHash, err = db.GetUserByName(username)
	u.name = username

	return u, err
}
