package Router

import (
	"encoding/json"
	"github.com/d97arkslayer/twitter-go/Middlewares"
	"github.com/d97arkslayer/twitter-go/Repositories"
	"net/http"
	"strconv"
)

/**
 * GetTweetsFollowers
 * Use to get tweets from followers
 */
func GetTweetsFollowers(w http.ResponseWriter, r *http.Request)  {
	if len(r.URL.Query().Get("page"))<1{
		http.Error(w, "Should sent the page parameter", http.StatusBadRequest)
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Should sent the page parameter like integer greater than 0", http.StatusBadRequest)
		return
	}
	tweets, status := Repositories.TweetsFollowers(Middlewares.IdUser, page)
	if status == false {
		http.Error(w,"Error getting the tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}