// upload.go

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const (
	staticDir        = "./static"
	maxImageFileSize = 20 << 20  // 20 MB maximum file size for images
	maxVideoFileSize = 100 << 20 // 100 MB maximum file size for videos
)

func uploadFile(w http.ResponseWriter, r *http.Request, fileType string) {
	// Set maximum file size based on the file type
	var maxFileSize int64
	switch fileType {
	case "images":
		maxFileSize = maxImageFileSize
	case "videos":
		maxFileSize = maxVideoFileSize
	default:
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}

	// Check the content length of the request
	if r.ContentLength > maxFileSize {
		http.Error(w, "File size exceeds the maximum allowed size", http.StatusBadRequest)
		return
	}

	// Get the file from the form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate a random UUID for the filename
	randomUUID := uuid.New()

	// Create the destination directory if it doesn't exist
	dir := fmt.Sprintf("%s/%s", staticDir, fileType)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}

	// Determine the file extension
	ext := filepath.Ext(handler.Filename)

	// Construct the destination filename
	destFilename := fmt.Sprintf("webwoods-cdn-%s%s", randomUUID, ext)

	// Create the destination file
	dst, err := os.Create(fmt.Sprintf("%s/%s", dir, destFilename))
	if err != nil {
		http.Error(w, "Unable to create destination file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the file content to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Unable to copy file content", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.Write([]byte("File uploaded successfully"))
}
