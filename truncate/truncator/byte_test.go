package truncator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/onozaty/filep/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewByteTruncator(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteBytes(
		t, d, "input", []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A})

	{
		output := filepath.Join(d, "output10")

		// ACT
		truncator, err := NewByteTruncator(10)
		require.NoError(t, err)
		err = truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A},
			test.ReadBytes(t, output))
	}
	{
		output := filepath.Join(d, "output9")

		// ACT
		truncator, err := NewByteTruncator(9)
		require.NoError(t, err)
		err = truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09},
			test.ReadBytes(t, output))
	}
	{
		output := filepath.Join(d, "output11")

		// ACT
		truncator, err := NewByteTruncator(11)
		require.NoError(t, err)
		err = truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A},
			test.ReadBytes(t, output))
	}
	{
		output := filepath.Join(d, "output0")

		// ACT
		truncator, err := NewByteTruncator(0)
		require.NoError(t, err)
		err = truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			[]byte{},
			test.ReadBytes(t, output))
	}
}

func TestNewByteTruncator_InputFileNotFound(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := filepath.Join(d, "xxxx")
	output := filepath.Join(d, "output")

	// ACT
	truncator, err := NewByteTruncator(10)
	require.NoError(t, err)
	err = truncator.Truncate(input, output)

	// ASSERT
	require.Error(t, err)
	pathErr := err.(*os.PathError)
	assert.Equal(t, input, pathErr.Path)
	assert.Equal(t, "open", pathErr.Op)
}

func TestNewByteTruncator_OutputFileNotFound(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteBytes(t, d, "input", []byte{})
	output := filepath.Join(d, "non", "output")

	// ACT
	truncator, err := NewByteTruncator(10)
	require.NoError(t, err)
	err = truncator.Truncate(input, output)

	// ASSERT
	require.Error(t, err)
	pathErr := err.(*os.PathError)
	assert.Equal(t, output, pathErr.Path)
	assert.Equal(t, "open", pathErr.Op)
}
