package schemas

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string   `json:"username"`
	PwdHash    string   `json:"-"`
	Email      string   `json:"email"`
	TelegramId int      `json:"telegramId"`
	Roles      []*Role  `json:"roles" gorm:"many2many:user_roles"`
	Groups     []*Group `json:"groups" gorm:"many2many:user_groups"`
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

func GetUserById(id uint) (*User, error) {
	var user *User
	result := db.First(user, id)
	return user, result.Error
}

func CheckPassword(user User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PwdHash), []byte(password))
	return err == nil
}
