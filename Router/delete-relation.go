package Router

import (
	"github.com/d97arkslayer/twitter-go/Middlewares"
	"github.com/d97arkslayer/twitter-go/Models"
	"github.com/d97arkslayer/twitter-go/Repositories"
	"github.com/gorilla/mux"
	"net/http"
)

/**
 * DeleteRelation
 * Http handler to delete a relation between users
 */
func DeleteRelation(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	if len(id) < 1 {
		http.Error(w, "You must provide a tweet id", http.StatusBadRequest)
		return
	}
	var relation Models.Relation
	relation.UserId = Middlewares.IdUser
	relation.UserRelationId = id
	status, err := Repositories.DeleteRelation(relation)
	if err != nil {
		if status == false {
			http.Error(w, "An error found trying to delete relation, Error: " + err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "An error found trying to delete relation " + err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Can not delete the relation", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}