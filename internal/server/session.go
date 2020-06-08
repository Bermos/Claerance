package server

import (
	"Claerance/internal/users"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handleSession(w http.ResponseWriter, r *http.Request, endpoint string) {
	log.Printf("Session - handling %s request for %s", r.Method, endpoint)

	switch endpoint {
	case "session":
		sessionBase(w, r)
	}
}

func sessionBase(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createSession(w, r)
	case "DELETE":
		destroySession(w, r)
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
	ip := r.Header.Get("X-Forwarded-Host")
	device := r.Header.Get("User-Agent")
	log.Println("login from:", login.Username, "with:", login.Password)

	if users.AuthWithCredentials(login.Username, login.Password, ip, device) {
		log.Println("Login successful")

		session, _ := store.Get(r, "claerance-session")
		// Set some session values.
		session.Values["authenticated"] = true
		session.Values["username"] = login.Username
		// Save it before we write to the response/return from the handler.
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else {
		log.Println("Login failed")
		http.Error(w, "Forbidden", http.StatusUnauthorized)
		//w.WriteHeader(http.StatusUnauthorized)
	}
}

func destroySession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "sso-session")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
