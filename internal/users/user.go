package users

import (
	"Claerance/internal/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

type User struct {
	gorm.Model
	Username   string `json:"username"`
	PwdHash    string `json:"-"`
	Email      string `json:"email"`
	TelegramId int    `json:"telegramId"`
}

func newUser(username string, pwdHash string, email string, telegramId int) User {
	return User{
		Username:   username,
		PwdHash:    pwdHash,
		Email:      email,
		TelegramId: telegramId,
	}
}

func CheckPassword(user User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PwdHash), []byte(password))
	return err == nil
}

func Setup() {
	db = database.GetDatabase()

	if err := db.AutoMigrate(&User{}); err != nil {
		log.Println("WARNING - Could not migrate db schema User")
	}

	CreateUser("Bermos", "test")
	CreateUser("Test", "test")
}

func CreateUser(username string, password string) {
	hashedPW, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	db.Create(&User{
		Username:   username,
		PwdHash:    string(hashedPW),
		Email:      "",
		TelegramId: 0,
	})
}

func GetUserById(userId uint) (User, error) {
	var user User
	result := db.First(&user, userId)
	return user, result.Error
}

func GetUserByName(username string) (User, error) {
	var user User
	result := db.First(&user, "username = ?", username)
	return user, result.Error
}

func GetAllUsers() ([]User, error) {
	var userList []User
	result := db.Find(&userList)
	return userList, result.Error
}

func DeleteUser(user User) bool {
	result := db.Delete(&user)
	return result.RowsAffected == 1
}

func DeleteUserById(userId uint) bool {
	result := db.Delete(&User{}, userId)
	return result.RowsAffected == 1
}

func UpdateUser(user User) error {
	result := db.Save(&user)
	return result.Error
}
