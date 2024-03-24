package main

import (
	"net/http"
)

func setupHandlers() {
	// Define the file server handler
	fs := http.FileServer(http.Dir("./static"))

	// Register the file server handler to serve static files
	http.Handle("/", fs)

	// images
	http.HandleFunc("/v1/api/upload/image", func(w http.ResponseWriter, r *http.Request) {
		uploadFile(w, r, "images")
	})
	http.HandleFunc("/v1/api/get/image", getImage)
	http.HandleFunc("/v1/api/get/images", getImages)

	// videos
	http.HandleFunc("/v1/api/upload/video", func(w http.ResponseWriter, r *http.Request) {
		uploadFile(w, r, "videos")
	})
	http.HandleFunc("/v1/api/get/video", getVideo)
	http.HandleFunc("/v1/api/get/videos", getVideos)
}
