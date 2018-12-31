package file

import (
	"bufio"
	"journ/crypto"
	"journ/file/journal"
	journOs "journ/file/os"
	params "journ/params/cli"
	"os"

	log "github.com/sirupsen/logrus"
)

// Encrypt file denoted by filePath
func Encrypt(filePath string, params params.CliParams) {
	if params.SkipEncryption {
		return
	}

	crypto.Encrypt(filePath, params.Key)
	journOs.DeleteFile(filePath)
}

// NewJournalPath returns a path pointing to new journal entry
func NewJournalPath() string {
	return journal.GetNewJournalPath()
}

// Decrypt file denoted by filePath
func Decrypt(filePath string, params params.CliParams) {
	contents := crypto.Decrypt(filePath, params.Key)
	normPathname := journal.NormalizePathname(filePath)

	file, err := os.Create(normPathname)

	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(file)
	_, err = w.WriteString(contents)

	if err != nil {
		log.Fatal(err)
	}

	w.Flush()

	log.Info("wrote decrypted contents to file: `" + normPathname + "'")
}
