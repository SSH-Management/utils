package utils

import (
	"io/fs"
	"os"
	"path/filepath"
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

func CreateDirectoryFromFile(path string, perm fs.FileMode) (string, error) {
	p, err := GetAbsolutePath(path)

	if err != nil {
		return "", err
	}

	directory := filepath.Dir(p)

	if err := os.MkdirAll(directory, perm); err != nil {
		return "", err
	}

	return p, nil
}

func CreateFile(path string, flags int, dirMode, mode fs.FileMode) (file *os.File, err error) {
	path, err = CreateDirectoryFromFile(path, fs.FileMode(dirMode)|fs.ModeDir)

	if err != nil {
		return nil, err
	}

	if _, err = os.Stat(path); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}

		file, err = os.Create(path)

		if err != nil {
			return nil, err
		}

		if err := file.Chmod(mode); err != nil {
			return nil, err
		}

		if err := file.Close(); err != nil {
			return nil, err
		}
	}

	file, err = os.OpenFile(path, flags, mode)

	return
}

func CreateLogFile(path string) (file *os.File, err error) {
	file, err = CreateFile(path, os.O_WRONLY|os.O_APPEND, 0744, fs.FileMode(0744)|os.ModeAppend)

	return
}

func FileExists(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		return false
	}

	return true
}

func CreateDirectory(path string, perm fs.FileMode) (string, error) {
	p, err := GetAbsolutePath(path)

	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(p, perm); err != nil {
		return "", err
	}

	return p, nil
}
