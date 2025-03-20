package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

func generateId(prefix string) string {
	buf := make([]byte, 12)
	rand.Read(buf)
	return fmt.Sprintf("%s_%s", prefix, hex.EncodeToString(buf))
}

func generateNow() int64 {
	return time.Now().UnixMilli()
}
