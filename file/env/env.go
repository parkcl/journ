package env

import (
	"os/user"

	log "github.com/sirupsen/logrus"
)

// GetUserHome directory as a string
func GetUserHome() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir
}
