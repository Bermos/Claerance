package server

import (
	"Claerance/internal/schemas"
	"github.com/gorilla/mux"
	"net/http"
)

func roleHandler(r *mux.Router) {
	r.HandleFunc("/create", createRole).Methods("POST")
	r.HandleFunc("/list", listRoles)
	r.HandleFunc("/{id:[0-9]+}", getRole).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", updateRole).Methods("PUT")
	r.HandleFunc("/{id:[0-9]+}", deleteRole).Methods("DELETE")
}

func createRole(w http.ResponseWriter, r *http.Request) {
	create(w, r, schemas.CreateRole)
}

func listRoles(w http.ResponseWriter, r *http.Request) {
	var roleList []schemas.Role
	readAll(w, r, &roleList)
}

func getRole(w http.ResponseWriter, r *http.Request) {
	read(w, r, &schemas.Role{})
}

func updateRole(w http.ResponseWriter, r *http.Request) {
	update(w, r, &schemas.Role{})
}

func deleteRole(w http.ResponseWriter, r *http.Request) {
	delete(w, r, &schemas.Role{})
}
