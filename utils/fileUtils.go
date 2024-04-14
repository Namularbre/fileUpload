package utils

import (
	"log"
	"os"
)

type FileData struct {
	Path string
	Name string
}

func GetAllFilesPaths(rootDir string) ([]FileData, error) {
	var subFilesPaths []FileData
	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		entryPath := rootDir + e.Name()
		fi, err := os.Stat(entryPath)
		if err != nil {
			return nil, err
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			subDirFilesPaths, err := GetAllFilesPaths(entryPath + "/")
			if err != nil {
				log.Fatal(err)
				return nil, err
			}
			subFilesPaths = append(subFilesPaths, subDirFilesPaths...)
		case mode.IsRegular():
			fd := FileData{
				Path: entryPath,
				Name: e.Name(),
			}

			subFilesPaths = append(subFilesPaths, fd)
		}
	}
	return subFilesPaths, nil
}

func (fd *FileData) ToString() string {
	return "Path: " + fd.Path + " Name: " + fd.Name
}
