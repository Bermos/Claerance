package server

import (
	"Claerance/internal/authentication"
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

	// Check if user is authenticated
	if authentication.TokenValid(r) {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Redirect(w, r, fmt.Sprintf("http://%s:%d/login", config.Cfg.Server.Host, config.Cfg.Server.Port), http.StatusSeeOther)
	}
}
