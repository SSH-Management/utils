package utils

import (
	"crypto/rand"
	"encoding/base64"
	"testing"
)

func TestRandomString(t *testing.T) {
	t.Parallel()

	l := 32

	str := RandomString(int32(l))

	if base64.RawURLEncoding.EncodedLen(l) != len(str) {
		t.Fatalf("Expected length: %d Given %d", l, len(str))
	}
}

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString(32)
	}
}

func random(b *testing.B) string {
	buffer := make([]byte, 32)

	n, err := rand.Read(buffer)

	if err != nil {
		b.Errorf("error while generating random buffer %v", err)
		return ""
	}

	if n != 32 {
		b.Errorf("expected length 32, given %d", n)
		return ""
	}

	return base64.RawURLEncoding.EncodeToString(buffer)
}

func BenchmarkCryptoRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random(b)
	}
}
