package server

import (
	"Claerance/internal/database"
	"Claerance/internal/schemas"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func userHandler(r *mux.Router) {
	r.HandleFunc("/create", createUser).Methods("POST")
	r.HandleFunc("/list", listUsers)
	r.HandleFunc("/{id:[0-9]+}", getUser).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", updateUser).Methods("PUT")
	r.HandleFunc("/{id:[0-9]+}", deleteUser).Methods("DELETE")
	r.HandleFunc("/", userBase)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	create(w, r, schemas.CreateUser)
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	var userList []schemas.User
	readAll(w, r, &userList)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	read(w, r, &schemas.User{})
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	update(w, r, &schemas.User{})
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	delete(w, r, &schemas.User{})
}

func userBase(w http.ResponseWriter, r *http.Request) {
	db := database.GetDatabase()
	var user schemas.User
	result := db.First(&user, 0) // TODO: GetUserID or make obsolete

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, encodeJson(user))
	}
}
