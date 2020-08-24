package Router

import (
	"github.com/d97arkslayer/twitter-go/Middlewares"
	"github.com/d97arkslayer/twitter-go/Models"
	"github.com/d97arkslayer/twitter-go/Repositories"
	"github.com/gorilla/mux"
	"net/http"
)

/**
 * HighRelation
 * Use to add relation between users
 */
func HighRelation(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	Id := vars["id"]
	if len(Id) < 1 {
		http.Error(w, "You must send the id parameter", http.StatusBadRequest)
		return
	}
	var relation Models.Relation
	relation.UserId = Middlewares.IdUser
	relation.UserRelationId = Id
	status, err := Repositories.StoreRelation(relation)
	if err != nil {
		http.Error(w, "An error ocurred trying insert relation in database " + err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Can't insert the relation in database", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}