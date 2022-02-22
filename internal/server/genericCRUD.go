package server

import (
	"Claerance/internal/authentication"
	"Claerance/internal/database"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

//create
func create(w http.ResponseWriter, r *http.Request, createFunc func(map[string]interface{})) {
	defer r.Body.Close()

	if !authentication.TokenValid(r) {
		w.WriteHeader(http.StatusForbidden)
	}

	var createRequest map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &createRequest)

	createFunc(createRequest)
}

//readAll list all of entities
func readAll(w http.ResponseWriter, r *http.Request, entityList interface{}) {
	db := database.GetDatabase()
	results := db.Find(entityList)

	if results.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, encodeJson(entityList))
}

//read get entity by id
func read(w http.ResponseWriter, r *http.Request, entity interface{}) {
	db := database.GetDatabase()
	result := db.First(entity, getIdFromRequest(r))

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, encodeJson(entity))
	}
}

func update(w http.ResponseWriter, r *http.Request, entity interface{}) {
	db := database.GetDatabase()

	// Get entity with ID from url param
	result := db.First(entity, getIdFromRequest(r))
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(GenericMsg{
			Msg: fmt.Sprintf("No %s with this id found.", reflect.TypeOf(entity)),
		})
		return
	}

	// Read body into empty interface
	var parsedUpdateRequest map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &parsedUpdateRequest)

	// Update found entity with parsed update request
	result = db.Model(entity).Updates(parsedUpdateRequest)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(GenericMsg{Msg: "No site with this id found."})
		return
	}
	w.WriteHeader(http.StatusOK)
}

//delete delete entity by it's id
func delete(w http.ResponseWriter, r *http.Request, entity interface{}) {
	db := database.GetDatabase()
	result := db.Delete(entity, getIdFromRequest(r))

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

//getIdFromRequest
func getIdFromRequest(r *http.Request) uint {
	vars := mux.Vars(r)
	ID, _ := strconv.ParseInt(vars["id"], 0, 64)

	return uint(ID)
}
