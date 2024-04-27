package main

import (
	"autoUpload/minio"
	"autoUpload/utils"
	"errors"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func parseArgs(args []string) (string, error) {
	filePath := args[0]
	info, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}

	if info.IsDir() {
		return filePath, nil
	}
	return "", errors.New("path must be a directory")
}

func getRootDir() string {
	argsWithoutProgram := os.Args[1:]

	if len(argsWithoutProgram) != 0 {
		rootDir, err := parseArgs(argsWithoutProgram)
		if err != nil {
			fmt.Printf("%v\n", err)
			panic("Cannot get information on upload directory.\n Try to add \"/\" at the end of the directory name")
		}
		return rootDir
	} else {
		return os.Getenv("ROOT_DIR_PATH")
	}
}

func main() {
	rootDir := getRootDir()

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
