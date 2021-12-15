package server

import (
	"Claerance/internal/config"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func authHandler(r *mux.Router) {
	r.HandleFunc("/", handleAuth)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	log.Printf("Auth - handling %s request", r.Method)

	session, _ := store.Get(r, "claerance-session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		log.Println("no auth:", r.Header.Get("X-Forwarded-Host"), "from:", r.Header.Get("X-Forwarded-For"))
		http.Redirect(w, r, fmt.Sprintf("http://%s:%d/login", config.Cfg.Server.Host, config.Cfg.Server.Port), http.StatusSeeOther) //TODO add host from config.Cfg.Server.Host
		return
	}

	log.Println("auth:   ", r.Header.Get("X-Forwarded-Host"), "from:", r.Header.Get("X-Forwarded-For"))
	w.WriteHeader(http.StatusOK)
}
