package content

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"webwoods.org/fileserver/internal/database"
)

type PresignedURLResponse struct {
	URL        string    `json:"url"`
	Expiration time.Time `json:"expiration"`
}

func GeneratePresignedUploadURL(w http.ResponseWriter, r *http.Request) {
	// Generate a unique folder ID for storing the uploaded files
	folderID := GenerateFolderID()

	// Create the folder if it doesn't exist
	folderPath := filepath.Join("./static", folderID, "images")
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		http.Error(w, "Failed to create folder", http.StatusInternalServerError)
		return
	}

	// Generate a unique filename
	fileName := generateFileName()

	// Calculate expiration time (10 seconds from now)
	expiration := time.Now().Add(3600 * time.Second)

	// Generate the URL for uploading
	uploadURL := fmt.Sprintf("/static/%s/images/%s", folderID, fileName)

	// Save the presigned URL and expiration date in MongoDB
	err = SavePresignedURL(uploadURL, expiration)
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

func SavePresignedURL(url string, expiration time.Time) error {
	mongoClient, err := database.GetMongoClient()

	if err != nil {
		return err
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

func IsURLExpired(url string) bool {
	fmt.Println("checking url validity: ", url)

	mongoClient, err := database.GetMongoClient()

	if err != nil {
		return true
	}

	collection := mongoClient.Database("reware").Collection("presigned_upload_urls")

	filter := bson.D{{Key: "url", Value: url}}
	opts := options.FindOne().SetSkip(0)

	var result PresignedURLResponse

	err = collection.FindOne(context.TODO(), filter, opts).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No documents found")
			return true
		} else {
			fmt.Println("Document found: ", result)
			return true
		}
	}

	var expirationStatus = time.Now().After(result.Expiration)
	fmt.Println(expirationStatus)

	return expirationStatus
}
