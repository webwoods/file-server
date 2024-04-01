// upload.go

package content

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func HandleFileUpload(w http.ResponseWriter, r *http.Request) {
	// Extract the filename from the URL path
	filename := filepath.Base(r.URL.Path)
	fmt.Println("Received file:", filename)

	if IsURLExpired(r.URL.Path) {
		http.Error(w, "Presigned URL is expired", http.StatusForbidden)
		return
	}

	// Define the directory where you want to save the files
	uploadDir := "./static"
	err := os.MkdirAll(uploadDir, 0755) // Create the directory if it doesn't exist
	if err != nil {
		http.Error(w, "Failed to create upload directory", http.StatusInternalServerError)
		return
	}
	fmt.Println("Upload directory exists")

	// Construct the file path
	relativePath := r.URL.Path[len("/static/"):] // Remove the "/static/" prefix
	filePath := filepath.Join(uploadDir, relativePath)

	fmt.Println("r.URL.Path: ", r.URL.Path)
	fmt.Println("Upload file path: ", filePath)

	// Create the file inside the upload directory
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Copy the data from the request body to the file
	_, err = io.Copy(file, r.Body)
	if err != nil {
		http.Error(w, "Failed to write file", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	fmt.Fprintf(w, "File %s uploaded successfully\n", filename)
}
