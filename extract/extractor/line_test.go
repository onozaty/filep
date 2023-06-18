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

func TestNewLineExtractor(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(
		t, d, "input", "1\n2\r\n3\n4\n\n6\r\n7xxxx\n8\n9\n10\n")

	{
		output := filepath.Join(d, "output1-10")

		// ACT
		extractor, _ := NewLineExtractor(1, 10, "UTF-8")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"1\n2\r\n3\n4\n\n6\r\n7xxxx\n8\n9\n10\n",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output2-9")

		// ACT
		extractor, _ := NewLineExtractor(2, 9, "UTF-8")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"2\r\n3\n4\n\n6\r\n7xxxx\n8\n9\n",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output1-11")

		// ACT
		extractor, _ := NewLineExtractor(1, 11, "UTF-8")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"1\n2\r\n3\n4\n\n6\r\n7xxxx\n8\n9\n10\n",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output11-12")

		// ACT
		extractor, _ := NewLineExtractor(11, 12, "UTF-8")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"",
			test.ReadString(t, output))
	}
}

func TestNewLineExtractor_改行無しで終了(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(
		t, d, "input", "\nあ2\n3\n4")

	{
		output := filepath.Join(d, "output1-4")

		// ACT
		extractor, _ := NewLineExtractor(1, 4, "UTF-8")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"\nあ2\n3\n4",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output2-3")

		// ACT
		extractor, _ := NewLineExtractor(2, 3, "UTF-8")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あ2\n3\n",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output2-2")

		// ACT
		extractor, _ := NewLineExtractor(2, 2, "utf-8")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あ2\n",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output1-1")

		// ACT
		extractor, _ := NewLineExtractor(1, 1, "UTF-8")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"\n",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output12-12")

		// ACT
		extractor, _ := NewLineExtractor(12, 12, "UTF-8")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"",
			test.ReadString(t, output))
	}
}

func TestNewLineExtractor_SJIS(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteBytes(
		t, d, "input.txt", test.StringToByte(t, "あい\nうえ\nお\nかき\nく\n", japanese.ShiftJIS))

	{
		output := filepath.Join(d, "output1-3")

		// ACT
		extractor, _ := NewLineExtractor(1, 3, "SJIS")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あい\nうえ\nお\n",
			test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS))
	}
	{
		output := filepath.Join(d, "output1-2")

		// ACT
		extractor, _ := NewLineExtractor(1, 2, "SJIS")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あい\nうえ\n",
			test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS))
	}
	{
		output := filepath.Join(d, "output1")

		// ACT
		extractor, _ := NewLineExtractor(1, 1, "sjis")
		err := extractor.Extract(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あい\n",
			test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS))
	}
}

func TestNewLineExtractor_InputFileNotFound(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := filepath.Join(d, "xxxx")
	output := filepath.Join(d, "output")

	// ACT
	extractor, _ := NewLineExtractor(1, 9, "UTF-8")
	err := extractor.Extract(input, output)

	// ASSERT
	require.Error(t, err)
	pathErr := err.(*os.PathError)
	assert.Equal(t, input, pathErr.Path)
	assert.Equal(t, "open", pathErr.Op)
}

func TestNewLineExtractor_OutputFileNotFound(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(t, d, "input", "")
	output := filepath.Join(d, "non", "output")

	// ACT
	extractor, _ := NewLineExtractor(1, 9, "UTF-8")
	err := extractor.Extract(input, output)

	// ASSERT
	require.Error(t, err)
	pathErr := err.(*os.PathError)
	assert.Equal(t, output, pathErr.Path)
	assert.Equal(t, "open", pathErr.Op)
}

func TestNewLineExtractor_InvalidEncoding(t *testing.T) {

	// ACT
	_, err := NewLineExtractor(1, 9, "utf")

	// ASSERT
	require.Error(t, err)
	assert.EqualError(t, err, "utf is invalid: htmlindex: invalid encoding name")
}

func TestNewLineExtractor_InvalidRange_Start(t *testing.T) {

	// ACT
	_, err := NewLineExtractor(0, 10, "utf-8")

	// ASSERT
	assert.EqualError(t, err, "invalid range: start = 0, end = 10")
}

func TestNewLineExtractor_InvalidRange_End(t *testing.T) {

	// ACT
	_, err := NewLineExtractor(10, 9, "utf-8")

	// ASSERT
	assert.EqualError(t, err, "invalid range: start = 10, end = 9")
}
