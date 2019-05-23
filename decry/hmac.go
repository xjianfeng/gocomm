package decry

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func HmacSha256(s string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
