package server

import (
	"Claerance/internal/users"
	userManager "Claerance/internal/users/manager"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserData struct {
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
}

func handleUser(w http.ResponseWriter, r *http.Request, endpoint string) {
	log.Printf("User - handling %s request for %s", r.Method, endpoint)

	switch endpoint {
	case "user":
		userBase(w, r)
	case "user/create":
		createUser(w, r)
	default:
		if strings.HasPrefix(endpoint, "user/") {
			switch r.Method {
			case "GET":
				getUser(w, r)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userManager.GetUserById(1)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user CreateUserRequest
	defer r.Body.Close()

	if !IsValid(r) {
		w.WriteHeader(http.StatusForbidden)
	}

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &user)

	userManager.CreateUser(user.Username, user.Password)
}

func userBase(w http.ResponseWriter, r *http.Request) {
	username := GetUsername(r)
	var user users.User
	user, _ = userManager.GetUserByName(username)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(UserData{Username: user.Username, CreatedAt: user.CreatedAt.String()})
}
