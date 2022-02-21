package server

import (
	"Claerance/internal/schemas"
	"github.com/gorilla/mux"
	"net/http"
)

func groupHandler(r *mux.Router) {
	r.HandleFunc("/create", createGroup).Methods("POST")
	r.HandleFunc("/list", listGroups).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", getGroup).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", updateGroup).Methods("PUT")
	r.HandleFunc("/{id:[0-9]+}", deleteGroup).Methods("DELETE")

}

func createGroup(w http.ResponseWriter, r *http.Request) {
	create(w, r, schemas.CreateGroup)
}

func listGroups(w http.ResponseWriter, r *http.Request) {
	var groupList []schemas.Group
	readAll(w, r, &groupList)
}

func getGroup(w http.ResponseWriter, r *http.Request) {
	read(w, r, &schemas.Group{})
}

func updateGroup(w http.ResponseWriter, r *http.Request) {
	update(w, r, &schemas.Group{})
}

func deleteGroup(w http.ResponseWriter, r *http.Request) {
	delete(w, r, &schemas.Group{})
}
