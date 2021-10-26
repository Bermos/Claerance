package users

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func (u User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PwdHash), []byte(password))
	return err == nil
}
