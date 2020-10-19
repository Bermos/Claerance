package users

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	PwdHash    string    `json:"-"`
	CreatedAt  time.Time `json:"createdAt"`
	Email      string    `json:"email"`
	TelegramId int       `json:"telegramId"`
}

func newUser(id int, username string, pwdHash string, createdAt time.Time, email string, telegramId int) User {
	return User{
		Id:         id,
		Username:   username,
		PwdHash:    pwdHash,
		CreatedAt:  createdAt,
		Email:      email,
		TelegramId: telegramId,
	}
}

func (u User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PwdHash), []byte(password))
	return err == nil
}
