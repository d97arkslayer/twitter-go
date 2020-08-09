package Middlewares

import (
	"github.com/d97arkslayer/twitter-go/Repositories"
	"github.com/d97arkslayer/twitter-go/Utils"
	"net/http"
)

/**
 * Email
 * user email
 */
var Email string
/**
 * IdUser
 * User id
 */
var IdUser string

/**
 * ValidateJWT
 * Use to check and validate JWT
 */
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		claims, err := Utils.ProcessToken(request.Header.Get("Authorization"))
		if err != nil {
			http.Error(writer, "Error in JWT " + err.Error(), http.StatusUnauthorized)
			return
		}
		_, exist, _ := Repositories.ExistUser(claims.Email)
		if exist==true {
			Email = claims.Email
			IdUser = claims.Id.Hex()
		}
		next.ServeHTTP(writer, request)
	}
}