package main

import (
	"log"
	"net/http"
	"os"

	"webwoods.org/fileserver/internal/api"
)

func main() {
	// Ensure that the static folder exists
	if _, err := os.Stat("./static"); os.IsNotExist(err) {
		err := os.Mkdir("./static", 0755)
		if err != nil {
			log.Fatalf("Failed to create static folder: %v", err)
		}
	}

	// Set up HTTP handlers
	api.SetupHandlers()

	// Start the server
	addr := ":8080"
	log.Printf("Server started at http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
