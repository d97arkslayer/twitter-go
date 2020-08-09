package Handlers

import (
	"github.com/d97arkslayer/twitter-go/Middlewares"
	"github.com/d97arkslayer/twitter-go/Router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)
/**
 * Handlers
 * Set de server port
 * And listen routes
 */
func Handlers(){
	router := mux.NewRouter()

	router.HandleFunc("/register", Middlewares.CheckDB(Router.Register)).Methods("POST")
	router.HandleFunc("/login", Middlewares.CheckDB(Router.Login)).Methods("POST")
	router.HandleFunc("/profile/{id}", Middlewares.CheckDB(Middlewares.ValidateJWT(Router.ShowProfile))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5050"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}