package Router

import (
	"github.com/d97arkslayer/twitter-go/Middlewares"
	"github.com/d97arkslayer/twitter-go/Repositories"
	"github.com/gorilla/mux"
	"net/http"
)

/**
 * DeleteTweet
 * Use to delete a tweet
 */
func DeleteTweet(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	if len(id) < 1 {
		http.Error(w, "You must provide a tweet id", http.StatusBadRequest)
		return
	}
	err := Repositories.DeleteTweet(id, Middlewares.IdUser)
	if err != nil {
		http.Error(w, "An error has occurred trying to delete the tweet from the database " + err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
