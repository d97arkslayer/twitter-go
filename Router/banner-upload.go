package Router

import (
	"github.com/d97arkslayer/twitter-go/Middlewares"
	"github.com/d97arkslayer/twitter-go/Models"
	"github.com/d97arkslayer/twitter-go/Repositories"
	"io"
	"net/http"
	"os"
	"strings"
)

/**
 * UploadBanner
 * Use this function to upload banner for a user
 */
func UploadBanner(w http.ResponseWriter, r *http.Request){
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var fileRoute string = "Uploads/Banners/" + Middlewares.IdUser + "." + extension
	f, err := os.OpenFile(fileRoute, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error uploading the image! " + err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copying the image! " + err.Error(), http.StatusBadRequest)
		return
	}
	var user Models.User
	var status bool
	user.Banner = Middlewares.IdUser + "." + extension
	status, err = Repositories.UpdateUser(user, Middlewares.IdUser)
	if err != nil || status == false {
		http.Error(w, "Error saving the banner in the Database " + err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}