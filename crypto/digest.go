package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
)

// Sha256Hex hash function
func Sha256Hex(data string) string {
	h256 := sha256.New()
	io.WriteString(h256, data)
	return fmt.Sprintf("%x", h256.Sum(nil))
}

// Md5Hex hash function
func Md5Hex(data string) string {
	m5 := md5.New()
	io.WriteString(m5, data)
	return fmt.Sprintf("%x", m5.Sum(nil))
}
