package base62

import (
	"bytes"
	"math/rand"
	"time"
)

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base     = 62
)

var (
	r = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
)

func Generate(length int) string {
	var buffer bytes.Buffer
	for i := 0; i < length; i++ {
		_ = buffer.WriteByte(alphabet[r.Intn(base)])
	}
	return buffer.String()
}
