package utils

import (
	"net/http"
	"os"
	"reflect"
	"unicode"
	"unsafe"
)

// #nosec G103
// UnsafeBytes returns a byte pointer without allocation
func UnsafeBytes(s string) (bs []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return
}

// #nosec G103
// UnsafeString returns a string pointer without allocation
func UnsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func IsSuccess(status int) bool {
	return status >= http.StatusOK && status < http.StatusMultipleChoices
}

func IsInt(s string) bool {
	if len(s) > 0 && s[0] == '0' {
		return false
	}

	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	return true
}

func Getenv(env string, def ...string) string {
	item := os.Getenv(env)

	if item == "" && len(def) > 0 {
		return def[0]
	}

	return item
}
