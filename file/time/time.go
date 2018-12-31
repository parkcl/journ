package time

import (
	"time"
)

func GetDateTimeFile() string {
	return time.Now().Format(time.RFC3339)
}
