package utils

import (
	"time"
)

func GetTimeMs() int64 {
	return time.Now().UnixNano() / 1000000
}
