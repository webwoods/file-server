package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Ensure that the static folder exists
	if _, err := os.Stat("./static"); os.IsNotExist(err) {
		err := os.Mkdir("./static", 0755)
		if err != nil {
			log.Fatalf("Failed to create static folder: %v", err)
		}
	}

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

	// Start the server
	addr := ":8080"
	log.Printf("Server started at http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
