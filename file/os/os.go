package os

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func CreateDirIfNotExists(path string) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}

	return path
}

func DeleteFile(filePath string) {
	if err := os.Remove(filePath); err != nil {
		log.Fatal(err)
	}

	log.Info("removed file " + filePath)
}
