package Router

import (
	"encoding/json"
	"github.com/d97arkslayer/twitter-go/Models"
	"github.com/d97arkslayer/twitter-go/Repositories"
	"github.com/d97arkslayer/twitter-go/Types"
	"github.com/d97arkslayer/twitter-go/Utils"
	"net/http"
	"time"
)

/**
 * Login
 * Do login
 */
func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type","application/json")
	var u Models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "User or Password INCORRECT " + err.Error(), http.StatusBadRequest)
		return
	}
	if len(u.Email) == 0{
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	user, exist := Repositories.Login(u.Email, u.Password)
	if exist == false {
		http.Error(w, "User or Password INCORRECT", http.StatusBadRequest)
		return
	}
	jwtKey, err := Utils.GenerateJWT(user)
	 if err != nil {
	 	http.Error(w, "An error occurred when the JWT was being generated " + err.Error(), http.StatusBadRequest)
		 return
	 }
	response := Types.LoginResponse{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})
}