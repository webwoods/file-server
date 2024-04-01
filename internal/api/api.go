package api

import (
	"net/http"

	"webwoods.org/fileserver/internal/content"
)

func SetupHandlers() {
	// Define the file server handler
	fs := http.FileServer(http.Dir("./static"))

	// Register the file server handler to serve static files
	http.Handle("/", fs)

	// presigned upload url
	http.HandleFunc("/v1/api/presigned/upload", content.GeneratePresignedUploadURL)

	// file upload
	http.HandleFunc("/static/", content.HandleFileUpload)
}
