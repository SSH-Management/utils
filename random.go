package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func RandomString(n int32) string {
	buffer := make([]byte, n)

	_, err := rand.Read(buffer)

	if err != nil {
		return ""
	}

	return base64.RawURLEncoding.EncodeToString(buffer)
}
