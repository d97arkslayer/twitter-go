package Router

import (
	"github.com/d97arkslayer/twitter-go/Repositories"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
)

/**
 * GetAvatar
 * Use this function to get user avatar
 */
func GetBanner(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	Id := vars["id"]
	if len(Id) < 1 {
		http.Error(w, "You must send the id parameter", http.StatusBadRequest)
		return
	}
	profile, err := Repositories.ShowProfile(Id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	Banner, err := os.Open("Uploads/Banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Banner not found", http.StatusNotFound)
		return
	}
	_, err = io.Copy(w, Banner)
	if err != nil {
		http.Error(w, "Error sending the banner image", http.StatusNotFound)
		return
	}
}