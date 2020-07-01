package manager

import (
	"Claerance/internal/database"
	"Claerance/internal/users"
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

func CreateUser(username string, password string) error {
	hashedPW, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return db.AddUser(username, string(hashedPW))
}

func GetUserById(userId int) (users.User, error) {
	u, err := db.GetUserById(userId)

	return u, err
}

func GetUserByName(username string) (users.User, error) {
	u, err := db.GetUserByName(username)

	return u, err
}
