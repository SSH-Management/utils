package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)


func TestCreateLogFile(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	path := "./test-logs/log.json"

	defer os.RemoveAll("./test-logs")

	file, err := CreateLogFile(path)

	assert.NoError(err)
	assert.NotNil(file)
	assert.FileExists(path)
}


func TestFileExists(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	t.Run("FileExists", func(t *testing.T) {
		file, err := os.Create("./test-exists.txt")

		assert.NoError(err)

		defer os.Remove("./test-exists.txt")
		defer file.Close()

		assert.True(FileExists("./test-exists.txt"))
	})

	t.Run("File_Does_Not_Exists", func(t *testing.T) {
		assert.False(FileExists("./test-does-not-exists.txt"))
	})
}


func TestCreateDirectory(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	defer os.RemoveAll("./test-dir")

	path, err := CreateDirectory("./test-dir", 0744)

	assert.NoError(err)
	abs, _ := filepath.Abs("./test-dir")

	assert.Equal(abs, path)
}
