package Router

import (
	"encoding/json"
	"github.com/d97arkslayer/twitter-go/Middlewares"
	"github.com/d97arkslayer/twitter-go/Models"
	"github.com/d97arkslayer/twitter-go/Repositories"
	"github.com/d97arkslayer/twitter-go/Types"
	"github.com/gorilla/mux"
	"net/http"
)

/**
 * Http handler to find a relation
 */
func FindRelation(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	if len(id) < 1 {
		http.Error(w, "You must provide a user relation id", http.StatusBadRequest)
		return
	}
	var relation Models.Relation
	relation.UserId = Middlewares.IdUser
	relation.UserRelationId = id
	var responseRelation Types.FindRelation

	status, err := Repositories.FindRelation(relation)
	if err != nil || status == false {
		responseRelation.Status = false
	} else {
		responseRelation.Status = true
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseRelation)
}