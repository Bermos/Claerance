package server

import (
	"Claerance/internal/entities/sites"
	"github.com/gorilla/mux"
	"net/http"
)

func siteHandler(r *mux.Router) {
	r.HandleFunc("/create", createSite).Methods("POST")
	r.HandleFunc("/list", listSites)
	r.HandleFunc("/{id:[0-9]+}", getSite).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", updateSite).Methods("PUT")
	r.HandleFunc("/{id:[0-9]+}", deleteSite).Methods("DELETE")
}

func createSite(w http.ResponseWriter, r *http.Request) {
	create(w, r, sites.CreateSite)
}

func listSites(w http.ResponseWriter, r *http.Request) {
	var siteList []sites.Site
	readAll(w, r, &siteList)
}

func getSite(w http.ResponseWriter, r *http.Request) {
	read(w, r, &sites.Site{})
}

func updateSite(w http.ResponseWriter, r *http.Request) {
	update(w, r, &sites.Site{})
}

func deleteSite(w http.ResponseWriter, r *http.Request) {
	delete(w, r, &sites.Site{})
}
