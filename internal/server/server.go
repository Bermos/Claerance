package server

import (
	"fmt"
	"github.com/gorilla/mux"
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
	r := mux.NewRouter()

	// Serve backend api
	apiEndpoint(r.PathPrefix("/api").Subrouter())

	// Serve static frontend page
	// TODO: fix abs path to prevent traversal
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := ioutil.ReadFile("web/public/" + r.URL.Path[1:])
		if err != nil {
			http.ServeFile(w, r, "web/public/index.html")
		} else {
			http.ServeContent(w, r, r.URL.Path[1:], time.Now(), strings.NewReader(string(file)))
		}
	})

	// Start server, this function is blocking
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func apiEndpoint(r *mux.Router) {
	authHandler(r.PathPrefix("/auth").Subrouter())
	userHandler(r.PathPrefix("/user").Subrouter())
	siteHandler(r.PathPrefix("/site").Subrouter())
	roleHandler(r.PathPrefix("/role").Subrouter())
	sessionHandler(r.PathPrefix("/session").Subrouter())

	// Catch faulty api requests
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Resource not found", http.StatusBadRequest)
	})
}
