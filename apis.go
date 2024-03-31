package main

import (
	"net/http"
)

func setupHandlers() {
	// Define the file server handler
	fs := http.FileServer(http.Dir("./static"))

	// Register the file server handler to serve static files
	http.Handle("/", fs)

	// presigned upload url
	http.HandleFunc("/v1/api/presigned/upload", generatePresignedUploadURL)

	// file upload
	http.HandleFunc("/static/", handleFileUpload)
}
