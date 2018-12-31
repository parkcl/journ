package cli

import (
	"fmt"
	"journ/file"
	params "journ/params/cli"
	"journ/vim"
	"os"

	flag "github.com/spf13/pflag"
	"golang.org/x/crypto/ssh/terminal"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	skipEncryption bool
	filePath       string
	passphrase     string
	decrypt        string
)

const (
	passphraseFlagName     = "passphrase"
	decryptFlagName        = "decrypt"
	skipEncryptionFlagName = "skip-encryption"
	stdinFd                = 0
)

var rootCmd = &cobra.Command{
	Use: "journ",
	Run: func(cmd *cobra.Command, args []string) {
		evalFlags(cmd.Flags())
	},
}

func evalFlags(flags *flag.FlagSet) {
	decryptFlag := flags.Lookup(decryptFlagName)

	if decryptFlag.Changed {
		file.Decrypt(decryptFlag.Value.String(), params.CliParams{Key: promptPassword()})
	} else {
		skipEncryptionFlag := flags.Lookup(skipEncryptionFlagName)

		if skipEncryptionFlag.Changed {
			vim.Open(params.CliParams{SkipEncryption: skipEncryption})
		} else {
			vim.Open(params.CliParams{SkipEncryption: skipEncryption, Key: promptPassword()})
		}
	}
}

func promptPassword() string {
	fmt.Print("Enter password: ")
	password, err := terminal.ReadPassword(stdinFd)

	if err != nil {
		log.Fatal(err)
	}

	return string(password)
}

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetReportCaller(true)

	rootCmd.Flags().BoolVarP(&skipEncryption, skipEncryptionFlagName, "s", false, "Skip encrypting file after")
	rootCmd.Flags().StringVarP(&decrypt, decryptFlagName, "d", "", "decrypt file")
}
