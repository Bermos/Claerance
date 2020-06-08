package server

import (
	"log"
	"net/http"
)

func handleAuth(w http.ResponseWriter, r *http.Request, endpoint string) {
	log.Printf("Auth - handling %s request for %s", r.Method, endpoint)

	session, _ := store.Get(r, "claerance-session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		log.Println("no auth:", r.Header.Get("X-Forwarded-Host"), "from:", r.Header.Get("X-Forwarded-For"))
		http.Redirect(w, r, "https://auth.bermos.dev", http.StatusSeeOther)
		return
	}

	log.Println("auth:   ", r.Header.Get("X-Forwarded-Host"), "from:", r.Header.Get("X-Forwarded-For"))
	w.WriteHeader(http.StatusOK)
}
