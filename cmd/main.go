package main

import (
	"Claerance/internal/users"
	"encoding/json"
	sess "github.com/gorilla/sessions"
	"io/ioutil"
	"log"
	"net/http"
)

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type GenericMsg struct {
	Msg string `json:"msg"`
}

var store = sess.NewCookieStore([]byte("asdjfadfasbfasdhfajf"))

func hauth(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "sso-session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		log.Println("no auth:", r.Header.Get("X-Forwarded-Host"), "from:", r.Header.Get("X-Forwarded-For"))
		http.Redirect(w, r, "https://auth.bermos.dev", http.StatusSeeOther)
		return
	}

	log.Println("auth:   ", r.Header.Get("X-Forwarded-Host"), "from:", r.Header.Get("X-Forwarded-For"))
	w.WriteHeader(http.StatusOK)
}

func sessions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(GenericMsg{Msg: "Method not allowed"})
	}

	session, _ := store.Get(r, "sso-session")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func login(w http.ResponseWriter, r *http.Request) {
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


		session, _ := store.Get(r, "sso-session")
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

func main() {
	fs := http.FileServer(http.Dir("webroot/public"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/auth", hauth)
	http.HandleFunc("/api/login", login)
	http.HandleFunc("/api/sessions", sessions)

	users.Setup()
	store.Options = &sess.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		Secure:   true,
		HttpOnly: true,
	}

	_ = http.ListenAndServe(":1401", nil)
}
