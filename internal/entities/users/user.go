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

func Setup() {
	db = database.GetDatabase()

	if err := db.AutoMigrate(&User{}); err != nil {
		log.Println("WARNING - Could not migrate db schema User")
	}

	CreateUser(map[string]interface{}{
		"username": "Bermos",
		"password": "test",
	})
	CreateUser(map[string]interface{}{
		"username": "Test",
		"password": "test",
	})
}

func CreateUser(request map[string]interface{}) {
	hashedPW, _ := bcrypt.GenerateFromPassword([]byte(request["password"].(string)), 12)
	db.Create(&User{
		Username:   request["username"].(string),
		PwdHash:    string(hashedPW),
		Email:      "",
		TelegramId: 0,
	})
}

func GetUserByName(username string) (User, error) {
	var user User
	result := db.First(&user, "username = ?", username)
	return user, result.Error
}

func CheckPassword(user User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PwdHash), []byte(password))
	return err == nil
}
