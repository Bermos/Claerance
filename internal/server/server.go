package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type GenericMsg struct {
	Msg string `json:"msg"`
}

// Start the server on the given port. This is blocking.
func Start(port int) {
	// Serve static frontend page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("web/public/" + r.URL.Path[1:])
		if err != nil {
			http.ServeFile(w, r, "web/public/index.html")
		} else {
			http.ServeContent(w, r, r.URL.Path[1:], time.Now(), strings.NewReader(string(file)))
		}
	})

	// Serve backend api
	http.HandleFunc("/api/", apiEndpoint)

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
