package authentication

import (
	"Claerance/internal/schemas"
	sess "github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var store *sess.CookieStore

const tokenName = "claerance-session"

func InitToken() {
	log.Info("Setting up sessions store")
	key := []byte("asdjfadfasbfasdhfajk")
	log.Info("Key length: %d", len(key))
	store = sess.NewCookieStore(key)

	store.Options = &sess.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		Secure:   true,
		HttpOnly: true,
	}

	log.Info("Sessions store setup")
}

func TokenValid(r *http.Request) bool {
	session, _ := store.Get(r, "claerance-session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		log.Debugf("no auth: %s from: %s", r.Header.Get("X-Forwarded-Host"), r.Header.Get("X-Forwarded-For"))
		return false
	} else {
		log.Debugf("auth:   %s from: %s", r.Header.Get("X-Forwarded-Host"), r.Header.Get("X-Forwarded-For"))
		return true
	}
}

func TokenUser(r *http.Request) *schemas.User {
	session, _ := store.Get(r, "claerance-session")
	user, err := schemas.GetUserById(session.Values["id"].(uint))
	if err != nil {
		return nil
	}

	return user
}

func TokenInvalidate(r *http.Request) *sess.Session {
	session, _ := store.Get(r, "claerance-session")

	// Revoke users authentication
	session.Values["authenticated"] = false
	return session
}

func TokenCreate(r *http.Request, user schemas.User) *sess.Session {
	session, _ := store.Get(r, "claerance-session")
	// Set some session values.
	session.Values["authenticated"] = true
	session.Values["username"] = user.Username
	session.Values["id"] = user.ID
	// Save it before writing to the response/return from the handler.
	return session
}
