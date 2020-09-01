package Router

import (
	"encoding/json"
	"github.com/d97arkslayer/twitter-go/Middlewares"
	"github.com/d97arkslayer/twitter-go/Repositories"
	"net/http"
	"strconv"
)

/**
 * IndexUsers
 * Use to List all users
 */
func IndexUsers(w http.ResponseWriter, r *http.Request){
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Should sent the page parameter, an integer greater tan 0", http.StatusBadRequest)
		return
	}
	pag := int64(pagTemp)
	users, status := Repositories.IndexUsers(Middlewares.IdUser, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error reading the users", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
