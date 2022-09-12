package truncator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/onozaty/filep/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/encoding/japanese"
)

func TestNewCharTruncator(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(
		t, d, "input", "あいうえお12345")

	{
		output := filepath.Join(d, "output10")

		// ACT
		truncator, _ := NewCharTruncator(10, "UTF-8")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あいうえお12345",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output9")

		// ACT
		truncator, _ := NewCharTruncator(9, "UTF-8")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あいうえお1234",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output11")

		// ACT
		truncator, _ := NewCharTruncator(11, "UTF-8")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あいうえお12345",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output0")
		truncator, _ := NewCharTruncator(0, "UTF-8")

		// ACT
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"",
			test.ReadString(t, output))
	}
}

func TestNewCharTruncator_SJIS(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteBytes(
		t, d, "input.txt", test.StringToByte(t, "あいうえおかきくけこ", japanese.ShiftJIS))

	{
		output := filepath.Join(d, "output10")

		// ACT
		truncator, _ := NewCharTruncator(10, "SJIS")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あいうえおかきくけこ",
			test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS))
	}
	{
		output := filepath.Join(d, "output9")
		truncator, _ := NewCharTruncator(9, "SJIS")

		// ACT
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あいうえおかきくけ",
			test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS))
	}
	{
		output := filepath.Join(d, "output11")

		// ACT
		truncator, _ := NewCharTruncator(11, "SJIS")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あいうえおかきくけこ",
			test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS))
	}
	{
		output := filepath.Join(d, "output0")
		truncator, _ := NewCharTruncator(0, "SJIS")

		// ACT
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"",
			test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS))
	}
}

func TestNewCharTruncator_InputFileNotFound(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := filepath.Join(d, "xxxx")
	output := filepath.Join(d, "output")

	// ACT
	truncator, _ := NewCharTruncator(9, "UTF-8")
	err := truncator.Truncate(input, output)

	// ASSERT
	require.Error(t, err)
	pathErr := err.(*os.PathError)
	assert.Equal(t, input, pathErr.Path)
	assert.Equal(t, "open", pathErr.Op)
}

func TestNewCharTruncator_OutputFileNotFound(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input", "")
	output := filepath.Join(d, "non", "output")

	// ACT
	truncator, _ := NewCharTruncator(9, "UTF-8")
	err := truncator.Truncate(input, output)

	// ASSERT
	require.Error(t, err)
	pathErr := err.(*os.PathError)
	assert.Equal(t, output, pathErr.Path)
	assert.Equal(t, "open", pathErr.Op)
}

func TestNewCharTruncator_InvalidCharset(t *testing.T) {

	// ACT
	_, err := NewCharTruncator(9, "utf")

	// ASSERT
	require.Error(t, err)
	assert.EqualError(t, err, "utf is invalid: htmlindex: invalid encoding name")
}
