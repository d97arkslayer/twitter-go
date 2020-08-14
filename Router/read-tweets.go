package Router

import (
	"encoding/json"
	"github.com/d97arkslayer/twitter-go/Repositories"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

/**
 * GetTweets
 * Handler to get all tweets
 */
func GetTweets(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	if len(id) < 1 {
		http.Error(w, "You should sent the param id", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You should sent the query param page", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "The page value must be great than Zero", http.StatusBadRequest)
		return
	}
	pag := int64(page)
	response, status := Repositories.GetTweets(id, pag)
	if status == false {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
