package hashgenerator

import (
	"crypto/sha256"
	"encoding/hex"
)

type sha256Generator struct {
	key string
}

func (h *sha256Generator) Generate(data []byte) string {
	t := sha256.New()
	t.Write(data)
	return hex.EncodeToString(t.Sum([]byte(h.key)))
}
