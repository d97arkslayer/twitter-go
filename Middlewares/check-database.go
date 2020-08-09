package Middlewares

import (
	"github.com/d97arkslayer/twitter-go/Database"
	"net/http"
)

/**
 * CheckDB
 * Check database connection middleware
 */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if Database.CheckConnection() == 0 {
			http.Error(w, "No database connection",500)
			return
		}
		next.ServeHTTP(w, r)
	}
}