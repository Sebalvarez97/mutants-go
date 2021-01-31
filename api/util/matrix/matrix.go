package matrix

import (
	"crypto/sha1"
	"encoding/hex"
)

func GenerateHashForStringArray(input []string) string {
	var sa string
	for _, v := range input {
		sa += v
	}
	h := sha1.New()
	h.Write([]byte(sa))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
