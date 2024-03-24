// image_handlers.go

package main

import (
	"fmt"
	"net/http"
)

func getImage(w http.ResponseWriter, r *http.Request) {
	// Serve image from the static/images directory
	http.ServeFile(w, r, fmt.Sprintf("%s/images/%s", staticDir, r.URL.Query().Get("filename")))
}

func getImages(w http.ResponseWriter, r *http.Request) {
	// Get list of image files from the static/images directory
	http.ServeFile(w, r, fmt.Sprintf("%s/images", staticDir))
}
