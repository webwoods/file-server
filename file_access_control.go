package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type PresignedURLResponse struct {
	URL        string    `json:"url"`
	Expiration time.Time `json:"expiration"`
}

func generatePresignedUploadURL(w http.ResponseWriter, r *http.Request) {
	// Generate a unique folder ID for storing the uploaded files
	folderID := generateFolderID()

	// Create the folder if it doesn't exist
	folderPath := filepath.Join("./static", "generated_"+folderID, "images")
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		http.Error(w, "Failed to create folder", http.StatusInternalServerError)
		return
	}

	// Generate a unique filename
	fileName := generateFileName()

	// Calculate expiration time (10 seconds from now)
	expiration := time.Now().Add(10 * time.Second)

	// Generate the URL for uploading
	uploadURL := fmt.Sprintf("/static/generated_%s/images/%s", folderID, fileName)

	// Save the presigned URL and expiration date in MongoDB
	err = savePresignedURL(uploadURL, expiration)
	if err != nil {
		http.Error(w, "Failed to save presigned URL", http.StatusInternalServerError)
		return
	}

	// Return the presigned URL and expiration time as JSON response
	response := PresignedURLResponse{
		URL:        uploadURL,
		Expiration: expiration,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func savePresignedURL(url string, expiration time.Time) error {
	mongoClient, err := GetMongoClient()

	if err != nil {
		println("error when getting client: %s", err.Error())
		return err
	}

	if mongoClient == nil {
		return errors.New("MongoDB client is nil")
	}

	collection := mongoClient.Database("reware").Collection("presigned_upload_urls")
	presignedUrl := PresignedURLResponse{
		URL:        url,
		Expiration: expiration,
	}

	result, err := collection.InsertOne(context.TODO(), presignedUrl)
	if err != nil {
		println(err.Error())
		return err
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return nil // No error occurred, return nil
}

func generateFolderID() string {
	// Generate a random folder ID (you can use UUID or any other method)
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func generateFileName() string {
	// Generate a random filename (you can use UUID or any other method)
	return fmt.Sprintf("%d.jpg", time.Now().UnixNano())
}

func isURLExpired(urlResponse PresignedURLResponse) bool {
	// Check if the presigned URL is expired
	return time.Now().After(urlResponse.Expiration)
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	// Extract the filename from the URL path
	filename := filepath.Base(r.URL.Path)
	fmt.Println("Received file:", filename)

	// Create a file with the same name as received from the URL
	file, err := os.Create(filename)
	if err != nil {
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
