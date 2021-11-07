package server

import (
	"Claerance/internal/entities/users"
	"encoding/json"
	"github.com/gorilla/mux"
	sess "github.com/gorilla/sessions"
	"io/ioutil"
	"log"
	"net/http"
)

var store *sess.CookieStore

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
	if IsValid(r) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(SessionInfo{UserId: GetUserId(r), Username: GetUsername(r), IsValid: true})
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
	log.Println("login from:", login.Username, "with:", login.Password)

	user, _ := users.GetUserByName(login.Username)
	if users.CheckPassword(user, login.Password) {
		log.Println("Login successful")

		session, _ := store.Get(r, "claerance-session")
		// Set some session values.
		session.Values["authenticated"] = true
		session.Values["username"] = login.Username
		session.Values["id"] = user.ID
		// Save it before writing to the response/return from the handler.
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	} else {
		log.Println("Login failed")
		http.Error(w, "Forbidden", http.StatusUnauthorized)
		//w.WriteHeader(http.StatusUnauthorized)
	}
}

func destroySession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "claerance-session")

	// Revoke users authentication
	session.Values["authenticated"] = false
	_ = session.Save(r, w)
}

func IsValid(r *http.Request) bool {
	session, _ := store.Get(r, "claerance-session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		log.Println("no auth:", r.Header.Get("X-Forwarded-Host"), "from:", r.Header.Get("X-Forwarded-For"))
		return false
	} else {
		log.Println("auth:   ", r.Header.Get("X-Forwarded-Host"), "from:", r.Header.Get("X-Forwarded-For"))
		return true
	}
}

func GetUsername(r *http.Request) string {
	session, _ := store.Get(r, "claerance-session")
	return session.Values["username"].(string)
}

func GetUserId(r *http.Request) uint {
	session, _ := store.Get(r, "claerance-session")
	return session.Values["id"].(uint)
}

func InitSessionStore() {
	log.Println("INFO - Setting up sessions store")
	key := []byte("asdjfadfasbfasdhfajk")
	log.Printf("INFO - Key length: %d", len(key))
	store = sess.NewCookieStore(key)

	store.Options = &sess.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		Secure:   true,
		HttpOnly: true,
	}

	log.Println("INFO - Sessions store setup")
}
