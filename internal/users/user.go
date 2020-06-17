package users

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id         int
	Username   string
	PwdHash    string
	CreatedAt  time.Time
	Email      string
	TelegramId int
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
