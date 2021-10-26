package server

import (
	"Claerance/internal/roles"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func roleHandler(r *mux.Router) {
	r.HandleFunc("/list", listRoles)
}

func listRoles(w http.ResponseWriter, r *http.Request) {
	siteList, err := roles.GetAllRoles()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, encodeJson(siteList))
}
