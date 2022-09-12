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

func TestNewLineTruncator(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(
		t, d, "input", "1\n2\r\n3\n4\n\n6\r\n7xxxx\n8\n9\n10\n")

	{
		output := filepath.Join(d, "output10")

		// ACT
		truncator, _ := NewLineTruncator(10, "UTF-8")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"1\n2\r\n3\n4\n\n6\r\n7xxxx\n8\n9\n10\n",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output9")

		// ACT
		truncator, _ := NewLineTruncator(9, "UTF-8")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"1\n2\r\n3\n4\n\n6\r\n7xxxx\n8\n9\n",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output11")

		// ACT
		truncator, _ := NewLineTruncator(11, "UTF-8")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"1\n2\r\n3\n4\n\n6\r\n7xxxx\n8\n9\n10\n",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output0")

		// ACT
		truncator, _ := NewLineTruncator(0, "UTF-8")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"",
			test.ReadString(t, output))
	}
}

func TestNewLineTruncator_改行無しで終了(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(
		t, d, "input", "\n2\n3\n4")

	{
		output := filepath.Join(d, "output4")

		// ACT
		truncator, _ := NewLineTruncator(4, "UTF-8")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"\n2\n3\n4",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output3")

		// ACT
		truncator, _ := NewLineTruncator(3, "UTF-8")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"\n2\n3\n",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output2")

		// ACT
		truncator, _ := NewLineTruncator(2, "UTF-8")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"\n2\n",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output1")

		// ACT
		truncator, _ := NewLineTruncator(1, "UTF-8")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"\n",
			test.ReadString(t, output))
	}
	{
		output := filepath.Join(d, "output0")

		// ACT
		truncator, _ := NewLineTruncator(0, "UTF-8")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"",
			test.ReadString(t, output))
	}
}

func TestNewLineTruncator_SJIS(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteBytes(
		t, d, "input.txt", test.StringToByte(t, "あい\nうえ\nお\nかき\nく\n", japanese.ShiftJIS))

	{
		output := filepath.Join(d, "output3")

		// ACT
		truncator, _ := NewLineTruncator(3, "SJIS")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あい\nうえ\nお\n",
			test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS))
	}
	{
		output := filepath.Join(d, "output2")
		truncator, _ := NewLineTruncator(2, "SJIS")

		// ACT
		err := truncator.Truncate(input, output)

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
		truncator, _ := NewLineTruncator(1, "SJIS")
		err := truncator.Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			"あい\n",
			test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS))
	}
	{
		output := filepath.Join(d, "output0")
		truncator, _ := NewLineTruncator(0, "SJIS")

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

func TestNewLineTruncator_InputFileNotFound(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := filepath.Join(d, "xxxx")
	output := filepath.Join(d, "output")

	// ACT
	truncator, _ := NewLineTruncator(9, "UTF-8")
	err := truncator.Truncate(input, output)

	// ASSERT
	require.Error(t, err)
	pathErr := err.(*os.PathError)
	assert.Equal(t, input, pathErr.Path)
	assert.Equal(t, "open", pathErr.Op)
}

func TestNewLineTruncator_OutputFileNotFound(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input", "")
	output := filepath.Join(d, "non", "output")

	// ACT
	truncator, _ := NewLineTruncator(9, "UTF-8")
	err := truncator.Truncate(input, output)

	// ASSERT
	require.Error(t, err)
	pathErr := err.(*os.PathError)
	assert.Equal(t, output, pathErr.Path)
	assert.Equal(t, "open", pathErr.Op)
}

func TestNewLineTruncator_InvalidEncoding(t *testing.T) {

	// ACT
	_, err := NewLineTruncator(9, "utf")

	// ASSERT
	require.Error(t, err)
	assert.EqualError(t, err, "utf is invalid: htmlindex: invalid encoding name")
}
