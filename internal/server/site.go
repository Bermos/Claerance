package server

import (
	"Claerance/internal/sites"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type CreateSiteRequest struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func siteHandler(r *mux.Router) {
	r.HandleFunc("/create", createSite).Methods("POST")
	r.HandleFunc("/list", listSites)
}

func createSite(w http.ResponseWriter, r *http.Request) {
	var site CreateSiteRequest
	defer r.Body.Close()

	if !IsValid(r) {
		w.WriteHeader(http.StatusForbidden)
	}

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &site)

	sites.CreateSite(site.Name, site.Url)
}

func listSites(w http.ResponseWriter, r *http.Request) {
	siteList, err := sites.GetAllSites()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, encodeJson(siteList))
}
