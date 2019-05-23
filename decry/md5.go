package decry

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Sum(signStr []byte) string {
	m := md5.New()
	m.Write(signStr)
	sign := m.Sum(nil)
	return hex.EncodeToString(sign)
}
