package Router

import (
	"encoding/json"
	"github.com/d97arkslayer/twitter-go/Middlewares"
	"github.com/d97arkslayer/twitter-go/Models"
	"github.com/d97arkslayer/twitter-go/Repositories"
	"github.com/d97arkslayer/twitter-go/Types"
	"net/http"
	"time"
)

/**
 * StoreTweet
 * The route to store a tweet
 */
func StoreTweet( w http.ResponseWriter, r *http.Request){
	var tweet Types.Tweet
	err := json.NewDecoder(r.Body).Decode(&tweet)
	register := Models.Tweet{
		UserId:  Middlewares.IdUser,
		Message: tweet.Message,
		Date: time.Now(),
	}
	_, status, err:= Repositories.StoreTweet(register)
	if err != nil {
		http.Error(w, "An error has occurred while inserting into the database"+ err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Could not insert into database", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
