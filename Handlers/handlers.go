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
	router.HandleFunc("/update-profile", Middlewares.CheckDB(Middlewares.ValidateJWT(Router.UpdateUser))).Methods("PUT", "PATCH")
	router.HandleFunc("/tweets", Middlewares.CheckDB(Middlewares.ValidateJWT(Router.StoreTweet))).Methods("POST")
	router.HandleFunc("/tweets/{id}", Middlewares.CheckDB(Middlewares.ValidateJWT(Router.GetTweets))).Methods("GET")
	router.HandleFunc("/tweets/{id}", Middlewares.CheckDB(Middlewares.ValidateJWT(Router.DeleteTweet))).Methods("DELETE")
	router.HandleFunc("/upload-avatar", Middlewares.CheckDB(Middlewares.ValidateJWT(Router.UploadAvatar))).Methods("POST")
	router.HandleFunc("/upload-banner", Middlewares.CheckDB(Middlewares.ValidateJWT(Router.UploadBanner))).Methods("POST")
	router.HandleFunc("/get-avatar/{id}", Middlewares.CheckDB(Router.GetAvatar)).Methods("GET")
	router.HandleFunc("/get-banner/{id}", Middlewares.CheckDB(Router.GetBanner)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5050"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}