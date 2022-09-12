package truncator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/onozaty/filep/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCharTruncator(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(
		t, d, "input", "あいうえお12345")

	{
		output := filepath.Join(d, "output10")
		truncator, _ := NewCharTruncator(10, "UTF-8")

		// ACT
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
		truncator, _ := NewCharTruncator(9, "UTF-8")

		// ACT
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
		truncator, _ := NewCharTruncator(11, "UTF-8")

		// ACT
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
