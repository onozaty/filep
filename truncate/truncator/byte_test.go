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
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteBytes(
		t, d, "input", []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A})

	{
		output := filepath.Join(d, "output10")

		// ACT
		err := NewByteTruncator(10).Truncate(input, output)

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
		err := NewByteTruncator(9).Truncate(input, output)

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
		err := NewByteTruncator(11).Truncate(input, output)

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
		err := NewByteTruncator(0).Truncate(input, output)

		// ASSERT
		require.NoError(t, err)
		assert.Equal(
			t,
			[]byte{},
			test.ReadBytes(t, output))
	}
}
