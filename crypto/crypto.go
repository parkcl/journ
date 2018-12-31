package crypto

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"

	log "github.com/sirupsen/logrus"
)

const (
	PgpMessage              = "PGP MESSAGE"
	extension               = ".pgp"
	incorrectPassphraseText = "passphrase incorrect"
)

// Encrypt the file represented by path filePath.
func Encrypt(filePath string, password string) {
	data, readErr := ioutil.ReadFile(filePath)
	outFile, createErr := os.Create(filePath + extension)

	if readErr != nil {
		log.Fatal(readErr)
	}
	if createErr != nil {
		log.Fatal(createErr)
	}
	log.Info("created output file for encrypted input file: " + outFile.Name())

	defer func() {
		if err := outFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	ciphertext, _ := armor.Encode(outFile, PgpMessage, nil)
	plaintext, _ := openpgp.SymmetricallyEncrypt(ciphertext, []byte(password), nil, nil)
	fmt.Fprintf(plaintext, string(data))
	plaintext.Close()
	ciphertext.Close()

	log.Info("successfully encrypted file")
}

// Decrypt pgp file using password
func Decrypt(filePath string, password string) string {
	log.Info("decrypting file `" + filePath + "'")

	encryptedFile, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	decbuf := bytes.NewBuffer(encryptedFile)
	decoded, err := armor.Decode(decbuf)

	if err != nil {
		log.Fatal(err)
	}

	timesCalled := 0

	md, err := openpgp.ReadMessage(decoded.Body, nil, func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		if timesCalled > 0 {
			return nil, errors.New(incorrectPassphraseText)
		}
		timesCalled++
		return []byte(password), nil
	}, nil)

	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(md.UnverifiedBody)

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
