package server

import (
	"Claerance/internal/users"
	userManager "Claerance/internal/users/manager"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func userHandler(r *mux.Router) {
	r.HandleFunc("/create", createUser).Methods("POST")
	r.HandleFunc("/list", listUsers)
	r.HandleFunc("/{id:[0-9]+}", getUser).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", deleteUser).Methods("DELETE")
	r.HandleFunc("/{id:[0-9]+}", updateUser).Methods("PUT")
	r.HandleFunc("/", userBase)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["id"], 0, 64)

	user, err := userManager.GetUserById(int(userId))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(GenericMsg{Msg: "No user with this id found."})
		return
	}

	if err := userManager.UpdateUser(user); err == nil {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["id"], 0, 64)
	success := userManager.DeleteUserById(int(userId))

	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["id"], 0, 64)
	user, err := userManager.GetUserById(int(userId))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, encodeJson(user))
	}
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	allUsers, err := userManager.GetAllUsers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, encodeJson(allUsers))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user CreateUserRequest
	defer r.Body.Close()

	if !IsValid(r) {
		w.WriteHeader(http.StatusForbidden)
	}

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &user)

	userManager.CreateUser(user.Username, user.Password)
}

func userBase(w http.ResponseWriter, r *http.Request) {
	username := GetUsername(r)
	var user users.User
	user, _ = userManager.GetUserByName(username)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, encodeJson(user))
}

func encodeJson(v interface{}) string {
	userJson, _ := json.Marshal(v)
	return string(userJson)
}
