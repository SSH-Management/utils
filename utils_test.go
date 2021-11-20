package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsSuccess(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	type Data struct {
		Value    int
		Expected bool
	}

	data := [400]Data{}

	for i := 0; i < 400; i++ {
		data[i] = Data{
			Value:    100 + i,
			Expected: 100+i >= 200 && 100+i < 300,
		}
	}

	for _, item := range data {
		assert.Equal(item.Expected, IsSuccess(item.Value))
	}
}

func TestUnsafeBytes(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	bytes := []byte("Hello World")

	unsafeBytes := UnsafeBytes("Hello World")

	assert.EqualValues(bytes, unsafeBytes)
}

func TestUnsafeString(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	bytes := []byte("Hello World")

	str := UnsafeString(bytes)

	assert.EqualValues("Hello World", str)
}

func TestGetenv(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	t.Run("DefaultValue", func(t *testing.T) {
		value := Getenv("HELLO_ENV")
		assert.Empty(value)

		value = Getenv("HELLO_ENV", "some_default_value")

		assert.NotEmpty(value)
		assert.Equal("some_default_value", value)
	})

	t.Run("WithEnvSet", func(t *testing.T) {
		os.Setenv("HELLO_ENV", "value")

		value := Getenv("HELLO_ENV")
		assert.NotEmpty(value)
		assert.Equal("value", value)

		value = Getenv("HELLO_ENV", "hello_world")
		assert.NotEmpty(value)
		assert.Equal("value", value)
	})
}

func TestIsInt(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	t.Run("IsInt", func(t *testing.T) {
		assert.True(IsInt("23445555"))
	})

	t.Run("NotAnInt", func(t *testing.T) {
		assert.False(IsInt("fjkhadskjdfhasjd"))
	})

	t.Run("CannotStartWith0", func(t *testing.T) {
		assert.False(IsInt("023355"))
	})
}
