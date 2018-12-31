package vim

import (
	"journ/file"
	params "journ/params/cli"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

const vim = "vim"

// Open vim
func Open(params params.CliParams) {
	newJournalPath := file.NewJournalPath()
	cmd := exec.Command(vim, newJournalPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	if params.ShouldEncrypt() {
		file.Encrypt(newJournalPath, params)
	}
}
