package truncator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/onozaty/filep/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLineTruncator(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(
		t, d, "input", "1\n2\r\n3\n4\n\n6\r\n7xxxx\n8\n9\n10\n")

	{
		output := filepath.Join(d, "output10")
		truncator, _ := NewLineTruncator(10, "UTF-8")

		// ACT
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
		truncator, _ := NewLineTruncator(9, "UTF-8")

		// ACT
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
		truncator, _ := NewLineTruncator(11, "UTF-8")

		// ACT
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
		truncator, _ := NewLineTruncator(0, "UTF-8")

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

func TestNewLineTruncator_改行無しで終了(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(
		t, d, "input", "\n2\n3\n4")

	{
		output := filepath.Join(d, "output4")
		truncator, _ := NewLineTruncator(4, "UTF-8")

		// ACT
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
		truncator, _ := NewLineTruncator(3, "UTF-8")

		// ACT
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
		truncator, _ := NewLineTruncator(2, "UTF-8")

		// ACT
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
		truncator, _ := NewLineTruncator(1, "UTF-8")
		// ACT
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
		truncator, _ := NewLineTruncator(0, "UTF-8")

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
