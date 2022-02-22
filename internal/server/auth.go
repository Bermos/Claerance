package server

import (
	"Claerance/internal/config"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func authHandler(r *mux.Router) {
	r.HandleFunc("/", handleAuth)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	log.Debugf("Auth - handling %s request", r.Method)

	session, _ := store.Get(r, "claerance-session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		log.Debugf("no auth: %s from: %s", r.Header.Get("X-Forwarded-Host"), r.Header.Get("X-Forwarded-For"))
		http.Redirect(w, r, fmt.Sprintf("http://%s:%d/login", config.Cfg.Server.Host, config.Cfg.Server.Port), http.StatusSeeOther)
		return
	}

	log.Debugf("auth:   %s from: %s", r.Header.Get("X-Forwarded-Host"), r.Header.Get("X-Forwarded-For"))
	w.WriteHeader(http.StatusOK)
}
