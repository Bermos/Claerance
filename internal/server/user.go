package server

import (
	"Claerance/internal/users"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handleUser(w http.ResponseWriter, r *http.Request, endpoint string) {
	log.Printf("User - handling %s request for %s", r.Method, endpoint)

	switch endpoint {
	case "user":
		userBase(w, r)
	case "user/create":
		createUser(w, r)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user CreateUserRequest
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &user)

	users.AddUser(user.Username, user.Password)
}

func userBase(w http.ResponseWriter, r *http.Request) {

}
