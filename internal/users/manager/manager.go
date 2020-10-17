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
	return db.GetUserById(userId)
}

func GetUserByName(username string) (users.User, error) {
	return db.GetUserByName(username)
}

func GetAllUsers() ([]users.User, error) {
	return db.GetAllUsers()
}

func DeleteUser(user users.User) bool {
	return DeleteUserById(user.Id)
}

func DeleteUserById(userId int) bool {
	return db.DeleteUserById(userId)
}

func UpdateUser(user users.User) (users.User, error) {
	return db.UpdateUser(user)
}
