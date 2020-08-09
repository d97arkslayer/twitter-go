package Router

import (
	"encoding/json"
	"github.com/d97arkslayer/twitter-go/Repositories"
	"github.com/gorilla/mux"
	"net/http"
)

/**
 * ShowProfile
 * Is the route to show profile
 */
func ShowProfile(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	Id := vars["id"]
	if len(Id) < 1 {
		http.Error(w, "You must send the id parameter", http.StatusBadRequest)
		return
	}
	profile, err := Repositories.ShowProfile(Id)
	if err != nil {
		http.Error(w, "An error occurred while searching the record" + err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}