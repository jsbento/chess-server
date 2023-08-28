package utils

import (
	"math/rand"
	"time"
)

func Rand64() uint64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return uint64(r.Int()) | (uint64(r.Int()) << 15) | (uint64(r.Int()) << 30) | (uint64(r.Int()) << 45) | ((uint64(r.Int()) & 0xf) << 60)
}
