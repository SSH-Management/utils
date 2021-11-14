package utils

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"unicode"
	"unsafe"

	"github.com/rs/zerolog/log"
)

func GetAbsolutePath(path string) (string, error) {
	var err error

	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)

		if err != nil {
			return "", err
		}

		return path, nil
	}

	return path, err
}

func CreatePath(path string, perm fs.FileMode) (string, error) {
	p, err := GetAbsolutePath(path)

	log.Debug().Str("path", p).Msg("Creating path")

	if err != nil {
		return "", err
	}

	directory := filepath.Dir(p)
	if err := os.MkdirAll(directory, perm); err != nil {
		return "", err
	}

	return p, nil
}

func CreateDirectory(path string, perm fs.FileMode) (string, error) {
	p, err := GetAbsolutePath(path)

	log.Debug().Str("path", p).Msg("Creating path")

	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(p, perm); err != nil {
		return "", err
	}

	return p, nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)

	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateLogFile(path string, perm fs.FileMode) (file *os.File, err error) {
	path, err = CreatePath(path, perm)

	if err != nil {
		return nil, err
	}

	if _, err = os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(path)

			if err != nil {
				return nil, err
			}

			if err := file.Chmod(perm); err != nil {
				return nil, err
			}

			if err := file.Close(); err != nil {
				return nil, err
			}

		} else {
			return nil, err
		}
	}

	file, err = os.OpenFile(path, os.O_WRONLY|os.O_APPEND, os.ModeAppend)

	return
}

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
