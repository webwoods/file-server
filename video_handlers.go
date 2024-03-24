// video_handlers.go

package main

import (
	"fmt"
	"net/http"
)

func getVideo(w http.ResponseWriter, r *http.Request) {
	// Serve video from the static/videos directory
	http.ServeFile(w, r, fmt.Sprintf("%s/videos/%s", staticDir, r.URL.Query().Get("filename")))
}

func getVideos(w http.ResponseWriter, r *http.Request) {
	// Get list of video files from the static/videos directory
	http.ServeFile(w, r, fmt.Sprintf("%s/videos", staticDir))
}
