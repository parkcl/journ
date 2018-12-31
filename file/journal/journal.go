package journal

import (
	"journ/file/env"
	journOs "journ/file/os"
	journTime "journ/file/time"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	RootJournDirectory      = "journals"
	NormalJournalFilePrefix = "journ-"
	JournalExtension        = ".journal"
	slash                   = "/"
	pgpExtension            = "pgp"
)

// GetNewJournalPath returns a string representing path of new journal entry
func GetNewJournalPath() string {
	var sb strings.Builder
	sb.WriteString(journOs.CreateDirIfNotExists(getJournalPath()))
	sb.WriteString(getJournalFile())

	return sb.String()
}

// NormalizePathname of path by removing any pgp extension
func NormalizePathname(path string) string {
	i := strings.LastIndex(path, pgpExtension)

	if i > -1 {
		return path[0 : i-1]
	}

	return path
}

func getJournalFile() string {
	var sb strings.Builder
	sb.WriteString(journTime.GetDateTimeFile())
	sb.WriteString(JournalExtension)

	return sb.String()
}

func getJournalPath() string {
	year, month, day := time.Now().Date()

	var sb strings.Builder
	sb.WriteString(env.GetUserHome())
	sb.WriteString(slash)
	sb.WriteString(RootJournDirectory)
	sb.WriteString(slash)
	sb.WriteString(strconv.Itoa(year))
	sb.WriteString(slash)
	sb.WriteString(month.String())
	sb.WriteString(slash)
	sb.WriteString(strconv.Itoa(day))
	sb.WriteString(slash)

	return filepath.FromSlash(sb.String())
}
