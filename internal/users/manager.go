package users

import (
	"Claerance/internal/database"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Login struct {
	ip         string
	device     string
	session    Session
	date       time.Time
	successful bool
}

type Session struct {
	id         int
	ip         string
	cookie     string
	device     string
	validUntil time.Time
}

type User struct {
	name    string
	pwdHash string
}

var (
	users []User
	db    database.Databaser
)

func Setup() {
	db = database.NewDatabase("sqlite3", "testUrl2")

	AddUser("bermos", "test")
	AddUser("test", "test")
}

func AddUser(username string, password string) {
	hashedPW, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	db.AddUser(username, string(hashedPW))
}

func GetUserById(userId int) (User, error) {
	var u User
	var err error
	u.name, err = db.GetUserById(userId)
	u.pwdHash = ""

	return u, err
}

func getUser(username string) (User, error) {
	var u User
	var err error
	u.pwdHash, err = db.GetUser(username)
	u.name = username

	return u, err
}

func getSession(username string) (User, error) {
	var foundUser User
	for _, user := range users {
		if user.name == username {
			return user, nil
		}
	}

	return foundUser, errors.New("no user found with that name")
}

func AuthWithCredentials(username string, password string, ip string, device string) bool {
	user, err := getUser(username)
	if err != nil {
		// user.logins = append(user.logins, Login{ip, device, Session{}, time.Now(), false})
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.pwdHash), []byte(password))
	success := err == nil

	//user.logins = append(user.logins, Login{ip, device, Session{}, time.Now(), success})

	return success
}

func loginUser(username string, password string, ip string, device string) bool {
	return false
}
