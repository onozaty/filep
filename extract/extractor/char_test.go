package extractor

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/onozaty/filep/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/encoding/japanese"
)

func TestNewCharExtractor(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(
		t, d, "input", "あいうえお12345")

	{
		output := filepath.Join(d, "output1-10")

		// ACT
		extractor, _ := NewCharExtractor(1, 10, "UTF-8")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あいうえお12345",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output2-9")

		// ACT
		extractor, _ := NewCharExtractor(2, 9, "utf-8")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"いうえお1234",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output1-11")

		// ACT
		extractor, _ := NewCharExtractor(1, 11, "UTF-8")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あいうえお12345",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output12-12")
		extractor, _ := NewCharExtractor(12, 12, "UTF-8")

		// ACT
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"",
			test.ReadString(t, output))
	}
}

func TestNewCharExtractor_SJIS(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteBytes(
		t, d, "input.txt", test.StringToByte(t, "あいうえおかきくけこ", japanese.ShiftJIS))

	{
		output := filepath.Join(d, "output5-10")

		// ACT
		extractor, _ := NewCharExtractor(5, 10, "SJIS")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"おかきくけこ",
			test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS))
	}
	{
		output := filepath.Join(d, "output1-9")
		extractor, _ := NewCharExtractor(1, 9, "sjis")

		// ACT
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あいうえおかきくけ",
			test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS))
	}
	{
		output := filepath.Join(d, "output10-11")

		// ACT
		extractor, _ := NewCharExtractor(10, 11, "SJIS")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"こ",
			test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS))
	}
	{
		output := filepath.Join(d, "output15-16")
		extractor, _ := NewCharExtractor(15, 16, "SJIS")

		// ACT
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"",
			test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS))
	}
}

func TestNewCharExtractor_InputFileNotFound(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := filepath.Join(d, "xxxx")
	output := filepath.Join(d, "output")

	// ACT
	extractor, _ := NewCharExtractor(1, 9, "UTF-8")
	err := extractor.Extract(input, output)

	// ASSERT
	require.Error(t, err)
	pathErr := err.(*os.PathError)
	assert.Equal(t, input, pathErr.Path)
	assert.Equal(t, "open", pathErr.Op)
}

func TestNewCharExtractor_OutputFileNotFound(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(t, d, "input", "")
	output := filepath.Join(d, "non", "output")

	// ACT
	extractor, _ := NewCharExtractor(1, 9, "UTF-8")
	err := extractor.Extract(input, output)

	// ASSERT
	require.Error(t, err)
	pathErr := err.(*os.PathError)
	assert.Equal(t, output, pathErr.Path)
	assert.Equal(t, "open", pathErr.Op)
}

func TestNewCharExtractor_InvalidEncoding(t *testing.T) {

	// ACT
	_, err := NewCharExtractor(1, 9, "utf")

	// ASSERT
	require.Error(t, err)
	assert.EqualError(t, err, "utf is invalid: htmlindex: invalid encoding name")
}

func TestNewCharExtractor_InvalidRange_Start(t *testing.T) {

	// ACT
	_, err := NewCharExtractor(0, 10, "utf-8")

	// ASSERT
	assert.EqualError(t, err, "invalid range: start = 0, end = 10")
}

func TestNewCharExtractor_InvalidRange_End(t *testing.T) {

	// ACT
	_, err := NewCharExtractor(10, 9, "utf-8")

	// ASSERT
	assert.EqualError(t, err, "invalid range: start = 10, end = 9")
}
