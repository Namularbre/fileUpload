package main

import (
	"autoUpload/minio"
	"autoUpload/utils"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	rootDir := os.Getenv("ROOT_DIR_PATH")

	filesData, err := utils.GetAllFilesPaths(rootDir)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Found %d files", len(filesData))

	for _, fileData := range filesData {
		minio.UploadFile(fileData.Name, fileData.Path)
	}
}
