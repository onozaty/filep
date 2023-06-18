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

func TestExtractCmd_File_Byte(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteBytes(t, d, "input", []byte{0x01, 0x02, 0x03, 0x04, 0x05})
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "2",
		"-e", "3",
		"-b",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	extracted := test.ReadBytes(t, output)
	assert.Equal(t, []byte{0x02, 0x03}, extracted)
}

func TestExtractCmd_File_Byte_Start(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteBytes(t, d, "input", []byte{0x01, 0x02, 0x03, 0x04, 0x05})
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "3",
		"-b",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	extracted := test.ReadBytes(t, output)
	assert.Equal(t, []byte{0x03, 0x04, 0x05}, extracted)
}

func TestExtractCmd_File_Byte_End(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteBytes(t, d, "input", []byte{0x01, 0x02, 0x03, 0x04, 0x05})
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-e", "1",
		"-b",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	extracted := test.ReadBytes(t, output)
	assert.Equal(t, []byte{0x01}, extracted)
}

func TestExtractCmd_File_Byte_Over(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteBytes(t, d, "input", []byte{0x01, 0x02, 0x03, 0x04, 0x05})
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "6",
		"-b",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	extracted := test.ReadBytes(t, output)
	assert.Equal(t, []byte{}, extracted)
}

func TestExtractCmd_File_Char(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(t, d, "input", "あいうえおかきくけこ")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "3",
		"-e", "9",
		"-c",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	extracted := test.ReadString(t, output)
	assert.Equal(t, "うえおかきくけ", extracted)
}

func TestExtractCmd_File_Char_Start(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(t, d, "input", "あいうえおかきくけこ")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "2",
		"-c",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	extracted := test.ReadString(t, output)
	assert.Equal(t, "いうえおかきくけこ", extracted)
}

func TestExtractCmd_File_Char_End(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(t, d, "input", "あいうえおかきくけこ")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-e", "2",
		"-c",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	extracted := test.ReadString(t, output)
	assert.Equal(t, "あい", extracted)
}

func TestExtractCmd_File_Char_Over(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(t, d, "input", "あいうえおかきくけこ")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "11",
		"-c",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	extracted := test.ReadString(t, output)
	assert.Equal(t, "", extracted)
}

func TestExtractCmd_File_Line(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(t, d, "input", "1あ\n2\r\n3\n4\n5\n")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "3",
		"-e", "4",
		"-l",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	extracted := test.ReadString(t, output)
	assert.Equal(t, "3\n4\n", extracted)
}

func TestExtractCmd_File_Line_Start(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(t, d, "input", "1あ\n2\r\n3\n4\n5\n")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "5",
		"-l",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	extracted := test.ReadString(t, output)
	assert.Equal(t, "5\n", extracted)
}

func TestExtractCmd_File_Line_End(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(t, d, "input", "1あ\n2\r\n3\n4\n5\n")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-e", "4",
		"-l",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	extracted := test.ReadString(t, output)
	assert.Equal(t, "1あ\n2\r\n3\n4\n", extracted)
}

func TestExtractCmd_File_Line_Over(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(t, d, "input", "1あ\n2\r\n3\n4\n5\n")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "6",
		"-l",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	extracted := test.ReadString(t, output)
	assert.Equal(t, "", extracted)
}

func TestExtractCmd_Dir_Byte(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateDir(t, d, "input")

	test.CreateFileWriteBytes(t, input, "input1", []byte{0x01})
	test.CreateFileWriteBytes(t, input, "input2", []byte{0x01, 0x02})
	test.CreateFileWriteBytes(t, input, "input3", []byte{0x01, 0x02, 0x03})

	output := test.CreateDir(t, d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "2",
		"-e", "2",
		"-b",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)
	{
		extracted := test.ReadBytes(t, filepath.Join(output, "input1"))
		assert.Equal(t, []byte{}, extracted)
	}
	{
		extracted := test.ReadBytes(t, filepath.Join(output, "input2"))
		assert.Equal(t, []byte{0x02}, extracted)
	}
	{
		extracted := test.ReadBytes(t, filepath.Join(output, "input3"))
		assert.Equal(t, []byte{0x02}, extracted)
	}
}

func TestExtractCmd_Dir_Char(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateDir(t, d, "input")

	test.CreateFileWriteString(t, input, "input1", "1234567890")
	test.CreateFileWriteString(t, input, "input2", "123456789")
	test.CreateFileWriteString(t, input, "input3", "12345678")

	output := test.CreateDir(t, d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "2",
		"-e", "9",
		"-c",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)
	{
		extracted := test.ReadString(t, filepath.Join(output, "input1"))
		assert.Equal(t, "23456789", extracted)
	}
	{
		extracted := test.ReadString(t, filepath.Join(output, "input2"))
		assert.Equal(t, "23456789", extracted)
	}
	{
		extracted := test.ReadString(t, filepath.Join(output, "input3"))
		assert.Equal(t, "2345678", extracted)
	}
}

func TestExtractCmd_Dir_Line(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateDir(t, d, "input")

	test.CreateFileWriteString(t, input, "input1", "1\n2\n3\n4\n5")
	test.CreateFileWriteString(t, input, "input2", "1\n2\n3\n4")
	test.CreateFileWriteString(t, input, "input3", "1\n2\n3")

	output := test.CreateDir(t, d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "2",
		"-e", "4",
		"-l",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)
	{
		extracted := test.ReadString(t, filepath.Join(output, "input1"))
		assert.Equal(t, "2\n3\n4\n", extracted)
	}
	{
		extracted := test.ReadString(t, filepath.Join(output, "input2"))
		assert.Equal(t, "2\n3\n4", extracted)
	}
	{
		extracted := test.ReadString(t, filepath.Join(output, "input3"))
		assert.Equal(t, "2\n3", extracted)
	}
}

func TestExtractCmd_Dir_CreateOutputDir(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateDir(t, d, "input")

	test.CreateFileWriteBytes(t, input, "input1", []byte{0x01, 0x02, 0x03, 0x04, 0x05})

	output := filepath.Join(d, "output") // 出力ディレクトリは存在しない状態

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "1",
		"-e", "1",
		"-b",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)
	extracted := test.ReadBytes(t, filepath.Join(output, "input1"))
	assert.Equal(t, []byte{0x01}, extracted)
}

func TestExtractCmd_Dir_Recursive(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

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
		"extract",
		"-i", input,
		"-s", "1",
		"-e", "1",
		"-c",
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

func TestExtractCmd_File_Char_SJIS(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteBytes(t, d, "input", test.StringToByte(t, "あいうえお", japanese.ShiftJIS))
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "1",
		"-e", "3",
		"-c",
		"-o", output,
		"--encoding", "sjis",
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	extracted := test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS)
	assert.Equal(t, "あいう", extracted)
}

func TestExtractCmd_File_Line_SJIS(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteBytes(t, d, "input", test.StringToByte(t, "あ\nい\r\nう\nえ\nお", japanese.ShiftJIS))
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "1",
		"-e", "2",
		"-l",
		"-o", output,
		"--encoding", "sjis",
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	extracted := test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS)
	assert.Equal(t, "あ\nい\r\n", extracted)
}

func TestExtractCmd_NoNumberSpecified(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(t, d, "input", "")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "1",
		"-e", "2",
		// -b、-c、-l のいずれも指定しない
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.EqualError(t, err, "specify one of the following: -b, -c, -l")
}

func TestExtractCmd_InvalidEncoding(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(t, d, "input", "")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "1",
		"-e", "2",
		"-l",
		"-o", output,
		"--encoding", "xxxx",
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.EqualError(t, err, "xxxx is invalid: htmlindex: invalid encoding name")
}

func TestExtractCmd_InputNotFound(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := filepath.Join(d, "input") // 存在しない
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "1",
		"-e", "2",
		"-l",
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

func TestExtractCmd_OutputNotFound(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateDir(t, d, "input")
	output := filepath.Join(d, "a", "b") // 親ディレクトリ自体が無い

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "1",
		"-e", "2",
		"-l",
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

func TestExtractCmd_InvalidStart(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(t, d, "input", "")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "0",
		"-e", "2",
		"-l",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.EqualError(t, err, "start must be greater than or equal to 1")
}

func TestExtractCmd_InvalidEnd(t *testing.T) {

	// ARRANGE
	d := t.TempDir()

	input := test.CreateFileWriteString(t, d, "input", "")
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"extract",
		"-i", input,
		"-s", "10",
		"-e", "9",
		"-l",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.EqualError(t, err, "end must be greater than or equal to start")
}
