package Router

import (
	"encoding/json"
	"github.com/d97arkslayer/twitter-go/Middlewares"
	"github.com/d97arkslayer/twitter-go/Models"
	"github.com/d97arkslayer/twitter-go/Repositories"
	"net/http"
)

/**
 * UpdateUser
 * Use to update user info
 */
func UpdateUser(w http.ResponseWriter, r *http.Request){
	var user Models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Wrong data " + err.Error(), http.StatusBadRequest)
		return
	}
	var status bool
	status, err = Repositories.UpdateUser(user, Middlewares.IdUser)
	if err != nil {
		http.Error(w, "An error occurred while updating user information" + err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Profile could not be updated " + err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}
