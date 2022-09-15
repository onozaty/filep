package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/onozaty/filep/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/encoding/japanese"
)

func TestTruncateCmd_File_Byte(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteBytes(t, d, "input", []byte{0x01, 0x02, 0x03, 0x04, 0x05})
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"truncate",
		"-i", input,
		"-b", "2",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	truncated := test.ReadBytes(t, output)
	assert.Equal(t, []byte{0x01, 0x02}, truncated)
}

func TestTruncateCmd_File_Char(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input", "1234567890")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"truncate",
		"-i", input,
		"-c", "5",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	truncated := test.ReadString(t, output)
	assert.Equal(t, "12345", truncated)
}

func TestTruncateCmd_File_Line(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input", "1\n2\r\n3\n4\n5\n")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"truncate",
		"-i", input,
		"-l", "2",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	truncated := test.ReadString(t, output)
	assert.Equal(t, "1\n2\r\n", truncated)
}

func TestTruncateCmd_Dir_Byte(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateDir(t, d, "input")

	test.CreateFileWriteBytes(t, input, "input1", []byte{0x01})
	test.CreateFileWriteBytes(t, input, "input2", []byte{0x01, 0x02})
	test.CreateFileWriteBytes(t, input, "input3", []byte{0x01, 0x02, 0x03})

	output := test.CreateDir(t, d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"truncate",
		"-i", input,
		"-b", "2",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)
	{
		truncated := test.ReadBytes(t, filepath.Join(output, "input1"))
		assert.Equal(t, []byte{0x01}, truncated)
	}
	{
		truncated := test.ReadBytes(t, filepath.Join(output, "input2"))
		assert.Equal(t, []byte{0x01, 0x02}, truncated)
	}
	{
		truncated := test.ReadBytes(t, filepath.Join(output, "input3"))
		assert.Equal(t, []byte{0x01, 0x02}, truncated)
	}
}

func TestTruncateCmd_Dir_Char(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateDir(t, d, "input")

	test.CreateFileWriteString(t, input, "input1", "1234567890")
	test.CreateFileWriteString(t, input, "input2", "123456789")
	test.CreateFileWriteString(t, input, "input3", "12345678")

	output := test.CreateDir(t, d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"truncate",
		"-i", input,
		"-c", "9",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)
	{
		truncated := test.ReadString(t, filepath.Join(output, "input1"))
		assert.Equal(t, "123456789", truncated)
	}
	{
		truncated := test.ReadString(t, filepath.Join(output, "input2"))
		assert.Equal(t, "123456789", truncated)
	}
	{
		truncated := test.ReadString(t, filepath.Join(output, "input3"))
		assert.Equal(t, "12345678", truncated)
	}
}

func TestTruncateCmd_Dir_Line(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateDir(t, d, "input")

	test.CreateFileWriteString(t, input, "input1", "1\n2\n3\n4\n5")
	test.CreateFileWriteString(t, input, "input2", "1\n2\n3\n4")
	test.CreateFileWriteString(t, input, "input3", "1\n2\n3")

	output := test.CreateDir(t, d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"truncate",
		"-i", input,
		"-l", "4",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)
	{
		truncated := test.ReadString(t, filepath.Join(output, "input1"))
		assert.Equal(t, "1\n2\n3\n4\n", truncated)
	}
	{
		truncated := test.ReadString(t, filepath.Join(output, "input2"))
		assert.Equal(t, "1\n2\n3\n4", truncated)
	}
	{
		truncated := test.ReadString(t, filepath.Join(output, "input3"))
		assert.Equal(t, "1\n2\n3", truncated)
	}
}

func TestTruncateCmd_Dir_CreateOutputDir(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateDir(t, d, "input")

	test.CreateFileWriteBytes(t, input, "input1", []byte{0x01, 0x02, 0x03, 0x04, 0x05})

	output := filepath.Join(d, "output") // 出力ディレクトリは存在しない状態

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"truncate",
		"-i", input,
		"-b", "1",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)
	truncated := test.ReadBytes(t, filepath.Join(output, "input1"))
	assert.Equal(t, []byte{0x01}, truncated)
}

func TestTruncateCmd_Dir_Recursive(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateDir(t, d, "input")
	test.CreateFileWriteString(t, input, "1.txt", "123")
	test.CreateFileWriteString(t, input, "2.txt", "")

	inputSub := test.CreateDir(t, input, "sub")
	test.CreateFileWriteString(t, inputSub, "3.txt", "abc")
	test.CreateFileWriteString(t, inputSub, "4.txt", "abcd")

	inputSubSub := test.CreateDir(t, inputSub, "sub")
	test.CreateFileWriteString(t, inputSubSub, "5.txt", "xyz")

	output := test.CreateDir(t, d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"truncate",
		"-i", input,
		"-c", "1",
		"-o", output,
		"--recursive",
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)
	{
		replaced := test.ReadString(t, filepath.Join(output, "1.txt"))
		assert.Equal(t, "1", replaced)
	}
	{
		replaced := test.ReadString(t, filepath.Join(output, "2.txt"))
		assert.Equal(t, "", replaced)
	}
	{
		replaced := test.ReadString(t, filepath.Join(output, "sub", "3.txt"))
		assert.Equal(t, "a", replaced)
	}
	{
		replaced := test.ReadString(t, filepath.Join(output, "sub", "4.txt"))
		assert.Equal(t, "a", replaced)
	}
	{
		replaced := test.ReadString(t, filepath.Join(output, "sub", "sub", "5.txt"))
		assert.Equal(t, "x", replaced)
	}
}

func TestTruncateCmd_File_Char_SJIS(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteBytes(t, d, "input", test.StringToByte(t, "あいうえお", japanese.ShiftJIS))
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"truncate",
		"-i", input,
		"-c", "3",
		"-o", output,
		"--encoding", "sjis",
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	truncated := test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS)
	assert.Equal(t, "あいう", truncated)
}

func TestTruncateCmd_File_Line_SJIS(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteBytes(t, d, "input", test.StringToByte(t, "あ\nい\r\nう\nえ\nお", japanese.ShiftJIS))
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"truncate",
		"-i", input,
		"-l", "2",
		"-o", output,
		"--encoding", "sjis",
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	truncated := test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS)
	assert.Equal(t, "あ\nい\r\n", truncated)
}

func TestTruncateCmd_NoNumberSpecified(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input", "")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"truncate",
		"-i", input,
		// -b、-c、-l のいずれも指定しない
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.EqualError(t, err, "no number is specified")
}

func TestTruncateCmd_InvalidEncoding(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input", "")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"truncate",
		"-i", input,
		"-l", "2",
		"-o", output,
		"--encoding", "xxxx",
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.EqualError(t, err, "xxxx is invalid: htmlindex: invalid encoding name")
}

func TestTruncateCmd_InputNotFound(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := filepath.Join(d, "input") // 存在しない
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"truncate",
		"-i", input,
		"-l", "2",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.Error(t, err)
	assert.True(t, os.IsNotExist(err))

	pathErr, ok := err.(*os.PathError)
	require.True(t, ok)
	assert.Equal(t, input, pathErr.Path)
}

func TestTruncateCmd_OutputNotFound(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateDir(t, d, "input")
	output := filepath.Join(d, "a", "b") // 親ディレクトリ自体が無い

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"truncate",
		"-i", input,
		"-l", "2",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.Error(t, err)
	assert.True(t, os.IsNotExist(err))

	pathErr, ok := err.(*os.PathError)
	require.True(t, ok)
	assert.Equal(t, output, pathErr.Path)
}
