package Routers

import (
	"encoding/json"
	"github.com/d97arkslayer/twitter-go/Database"
	"github.com/d97arkslayer/twitter-go/Models"
	"net/http"
)

/**
 * Register
 * This function register a new user
 */
func Register(w http.ResponseWriter, r *http.Request){
	var t Models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Bad Request, please review all fields " + err.Error(), http.StatusBadRequest)
		return
	}
	if len(t.Email)==0 {
		http.Error(w, "The user email is required", http.StatusBadRequest)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "The password must be contain at least 6 characters", http.StatusBadRequest)
		return
	}
	_,exists,_ := Database.ExistUser(t.Email)
	if exists == true {
		http.Error(w, "The user already exist", http.StatusBadRequest)
		return
	}
	_, status, err := Database.InsertUser(t)
	if err != nil {
		http.Error(w, "An error has occurred when trying to insert the information in the database " + err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "The record could not be inserted into the database", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}