package hashkit

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
