package content

import (
	"fmt"
	"time"
)

func GenerateFolderID() string {
	// Generate a random folder ID (you can use UUID or any other method)
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func generateFileName() string {
	// Generate a random filename (you can use UUID or any other method)
	return fmt.Sprintf("%d.jpg", time.Now().UnixNano())
}
