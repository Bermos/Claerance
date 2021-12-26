package server

import (
	"Claerance/internal/database"
	"Claerance/internal/schemas"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func siteHandler(r *mux.Router) {
	r.HandleFunc("/create", createSite).Methods("POST")
	r.HandleFunc("/list", listSites)
	r.HandleFunc("/{id:[0-9]+}", getSite).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", updateSite).Methods("PUT")
	r.HandleFunc("/{id:[0-9]+}", deleteSite).Methods("DELETE")
	r.HandleFunc("/{id:[0-9]+}/authorized", getSiteAuthorized).Methods("GET")
}

func createSite(w http.ResponseWriter, r *http.Request) {
	create(w, r, schemas.CreateSite)
}

func listSites(w http.ResponseWriter, r *http.Request) {
	var siteList []schemas.Site
	readAll(w, r, &siteList)
}

func getSite(w http.ResponseWriter, r *http.Request) {
	read(w, r, &schemas.Site{})
}

func updateSite(w http.ResponseWriter, r *http.Request) {
	update(w, r, &schemas.Site{})
}

func deleteSite(w http.ResponseWriter, r *http.Request) {
	delete(w, r, &schemas.Site{})
}

func getSiteAuthorized(w http.ResponseWriter, r *http.Request) {
	var authorizedEntities schemas.SiteAuthorizedEntities
	var site schemas.Site
	db := database.GetDatabase()

	results := db.Find(&site)
	if results.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	authorizedEntities.Users = site.AuthorizedUsers
	authorizedEntities.Groups = site.AuthorizedGroups
	fmt.Fprintf(w, encodeJson(authorizedEntities))
}
