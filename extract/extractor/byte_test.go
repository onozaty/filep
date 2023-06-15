package extractor

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/onozaty/filep/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewByteExtractor(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteBytes(
		t, d, "input", []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A})

	{
		output := filepath.Join(d, "output1-10")

		// ACT
		extractor, err := NewByteExtractor(1, 10)
		require.NoError(t, err)
		err = extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A},
			test.ReadBytes(t, output))
	}
	{
		output := filepath.Join(d, "output2-9")

		// ACT
		extractor, err := NewByteExtractor(2, 9)
		require.NoError(t, err)
		err = extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			[]byte{0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09},
			test.ReadBytes(t, output))
	}
	{
		output := filepath.Join(d, "output1-11")

		// ACT
		extractor, err := NewByteExtractor(1, 11)
		require.NoError(t, err)
		err = extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A},
			test.ReadBytes(t, output))
	}
	{
		output := filepath.Join(d, "output12-13")

		// ACT
		extractor, err := NewByteExtractor(12, 13)
		require.NoError(t, err)
		err = extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			[]byte{},
			test.ReadBytes(t, output))
	}
}

func TestNewByteExtractor_InputFileNotFound(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := filepath.Join(d, "xxxx")
	output := filepath.Join(d, "output")

	// ACT
	extractor, err := NewByteExtractor(1, 10)
	require.NoError(t, err)
	err = extractor.Extract(input, output)

	// ASSERT
	require.Error(t, err)
	pathErr := err.(*os.PathError)
	assert.Equal(t, input, pathErr.Path)
	assert.Equal(t, "open", pathErr.Op)
}

func TestNewByteExtractor_OutputFileNotFound(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteBytes(t, d, "input", []byte{})
	output := filepath.Join(d, "non", "output")

	// ACT
	extractor, err := NewByteExtractor(1, 10)
	require.NoError(t, err)
	err = extractor.Extract(input, output)

	// ASSERT
	require.Error(t, err)
	pathErr := err.(*os.PathError)
	assert.Equal(t, output, pathErr.Path)
	assert.Equal(t, "open", pathErr.Op)
}

func TestNewByteExtractor_InvalidRange_Start(t *testing.T) {

	// ACT
	_, err := NewByteExtractor(0, 10)

	// ASSERT
	assert.EqualError(t, err, "invalid range: start = 0, end = 10")
}

func TestNewByteExtractor_InvalidRange_End(t *testing.T) {

	// ACT
	_, err := NewByteExtractor(10, 9)

	// ASSERT
	assert.EqualError(t, err, "invalid range: start = 10, end = 9")
}
