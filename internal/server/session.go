package server

import (
	"Claerance/internal/authentication"
	"Claerance/internal/schemas"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SessionInfo struct {
	UserId   uint   `json:"userId,omitempty"`
	Username string `json:"username,omitempty"`
	IsValid  bool   `json:"isValid"`
}

func sessionHandler(r *mux.Router) {
	r.HandleFunc("/", createSession).Methods("POST")
	r.HandleFunc("/", destroySession).Methods("DELETE")
	r.HandleFunc("/", getSessionValid).Methods("GET")
}

func getSessionValid(w http.ResponseWriter, r *http.Request) {
	if user := authentication.TokenUser(r); user != nil {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(SessionInfo{UserId: user.ID, Username: user.Username, IsValid: true})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(SessionInfo{IsValid: false})
	}
}

func createSession(w http.ResponseWriter, r *http.Request) {
	var login LoginRequest
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &login)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(GenericMsg{Msg: "Not proper JSON"})
		return
	}
	/*ip := r.Header.Get("X-Forwarded-Host")
	device := r.Header.Get("User-Agent")*/
	log.Debugf("login from: %s with: %s", login.Username, login.Password)

	user, _ := schemas.GetUserByName(login.Username)
	if schemas.CheckPassword(user, login.Password) {
		log.Debug("Login successful")

		err := authentication.TokenCreate(r, user).Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	} else {
		log.Debug("Login failed")
		http.Error(w, "Forbidden", http.StatusUnauthorized)
		//w.WriteHeader(http.StatusUnauthorized)
	}
}

func destroySession(w http.ResponseWriter, r *http.Request) {
	_ = authentication.TokenInvalidate(r).Save(r, w)
}
