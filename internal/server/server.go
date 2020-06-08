package server

import (
	"encoding/json"
	"fmt"
	sess "github.com/gorilla/sessions"
	"log"
	"net/http"
	"strings"
)

type GenericMsg struct {
	Msg string `json:"msg"`
}

var store = sess.NewCookieStore([]byte("asdjfadfasbfasdhfajf"))

// Start the server on the given port. This is blocking.
func Start(port int) {
	// Serve static frontend page
	fs := http.FileServer(http.Dir("webroot/public"))
	http.Handle("/", http.StripPrefix("/", fs))

	http.HandleFunc("/api/", apiEndpoint)

	store.Options = &sess.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		Secure:   true,
		HttpOnly: true,
	}

	// Start server, this function is blocking
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func apiEndpoint(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.RequestURI, "/")

	if len(urlParts) < 4 {
		http.Error(w, "Resource not found", http.StatusBadRequest)
		return
	}

	if urlParts[2] != "v1" {
		_ = json.NewEncoder(w).Encode(GenericMsg{Msg: "Unknown API version"})
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	endpoint := strings.Join(urlParts[3:], "/")
	switch urlParts[3] {
	case "user":
		handleUser(w, r, endpoint)
	case "auth":
		handleAuth(w, r, endpoint)
	case "session":
		handleSession(w, r, endpoint)
	}
}
