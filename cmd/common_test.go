package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/text/encoding"
)

func createFileWriteBytes(t *testing.T, dir string, name string, content []byte) string {

	file, err := os.Create(filepath.Join(dir, name))
	require.NoError(t, err)

	_, err = file.Write(content)
	require.NoError(t, err)

	err = file.Close()
	require.NoError(t, err)

	return file.Name()
}

func createFileWriteString(t *testing.T, dir string, name string, content string) string {

	return createFileWriteBytes(t, dir, name, []byte(content))
}

func createTempDir(t *testing.T) string {

	tempDir, err := os.MkdirTemp("", "filep")
	require.NoError(t, err)

	return tempDir
}

func createDir(t *testing.T, parent string, name string) string {

	dir := filepath.Join(parent, name)
	err := os.Mkdir(dir, os.ModePerm)
	require.NoError(t, err)

	return dir
}

func readBytes(t *testing.T, name string) []byte {

	bo, err := os.ReadFile(name)
	require.NoError(t, err)

	return bo
}

func readString(t *testing.T, name string) string {

	bo := readBytes(t, name)
	return string(bo)
}

func byteToString(t *testing.T, bytes []byte, enc encoding.Encoding) string {
	decoded, err := enc.NewDecoder().Bytes(bytes)
	require.NoError(t, err)

	return string(decoded)
}

func stringToByte(t *testing.T, str string, enc encoding.Encoding) []byte {
	encoded, err := enc.NewEncoder().Bytes([]byte(str))
	require.NoError(t, err)

	return encoded
}
