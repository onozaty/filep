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

func TestReplaceCmd_File_Regex(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input.txt", "abc\nabc\naa")
	output := filepath.Join(d, "output.txt")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-r", "a$",
		"-t", "x",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	replaced := test.ReadString(t, output)
	assert.Equal(t, "abc\nabc\nax", replaced)
}

func TestReplaceCmd_File_String(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input.txt", "aa.ab.ac.ad.a.b.c.d")
	output := filepath.Join(d, "output.txt")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-s", "a.",
		"-t", "xx",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	replaced := test.ReadString(t, output)
	assert.Equal(t, "axxab.ac.ad.xxb.c.d", replaced)
}

func TestReplaceCmd_Dir_Regex(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateDir(t, d, "input")

	test.CreateFileWriteString(t, input, "input1.txt", "abc\nabc\naa")
	test.CreateFileWriteString(t, input, "input2.txt", "a")
	test.CreateFileWriteString(t, input, "input3.txt", "ax")

	output := test.CreateDir(t, d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-r", "a$",
		"-t", "x",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)
	{
		replaced := test.ReadString(t, filepath.Join(output, "input1.txt"))
		assert.Equal(t, "abc\nabc\nax", replaced)
	}
	{
		replaced := test.ReadString(t, filepath.Join(output, "input2.txt"))
		assert.Equal(t, "x", replaced)
	}
	{
		replaced := test.ReadString(t, filepath.Join(output, "input3.txt"))
		assert.Equal(t, "ax", replaced)
	}
}

func TestReplaceCmd_File_Regex_Japanese(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input.txt", "あいうえおかきくけこ")
	output := filepath.Join(d, "output.txt")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-r", "あ.{4}",
		"-t", "",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	replaced := test.ReadString(t, output)
	assert.Equal(t, "かきくけこ", replaced)
}

func TestReplaceCmd_Dir_String(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateDir(t, d, "input")

	test.CreateFileWriteString(t, input, "input1.txt", "abc\na.c\naa")
	test.CreateFileWriteString(t, input, "input2.txt", "")
	test.CreateFileWriteString(t, input, "input3.txt", "a.c")

	output := test.CreateDir(t, d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-s", "a.c",
		"-t", "",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)
	{
		replaced := test.ReadString(t, filepath.Join(output, "input1.txt"))
		assert.Equal(t, "abc\n\naa", replaced)
	}
	{
		replaced := test.ReadString(t, filepath.Join(output, "input2.txt"))
		assert.Equal(t, "", replaced)
	}
	{
		replaced := test.ReadString(t, filepath.Join(output, "input3.txt"))
		assert.Equal(t, "", replaced)
	}
}

func TestReplaceCmd_Dir_CreateOutputDir(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateDir(t, d, "input")

	test.CreateFileWriteString(t, input, "input1.txt", "abc\nabc\nabc")

	output := filepath.Join(d, "output") // 出力ディレクトリは存在しない状態

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-r", "(?m)c$",
		"-t", "x",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	replaced := test.ReadString(t, filepath.Join(output, "input1.txt"))
	assert.Equal(t, "abx\nabx\nabx", replaced)
}

func TestReplaceCmd_Dir_Recursive(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateDir(t, d, "input")
	test.CreateFileWriteString(t, input, "1.txt", "abc")
	test.CreateFileWriteString(t, input, "2.txt", "")

	inputSub := test.CreateDir(t, input, "sub")
	test.CreateFileWriteString(t, inputSub, "3.txt", "cat")
	test.CreateFileWriteString(t, inputSub, "4.txt", "aaa")

	inputSubSub := test.CreateDir(t, inputSub, "sub")
	test.CreateFileWriteString(t, inputSubSub, "5.txt", "a")

	output := test.CreateDir(t, d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-s", "a",
		"-t", "",
		"--recursive",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)
	{
		replaced := test.ReadString(t, filepath.Join(output, "1.txt"))
		assert.Equal(t, "bc", replaced)
	}
	{
		replaced := test.ReadString(t, filepath.Join(output, "2.txt"))
		assert.Equal(t, "", replaced)
	}
	{
		replaced := test.ReadString(t, filepath.Join(output, "sub", "3.txt"))
		assert.Equal(t, "ct", replaced)
	}
	{
		replaced := test.ReadString(t, filepath.Join(output, "sub", "4.txt"))
		assert.Equal(t, "", replaced)
	}
	{
		replaced := test.ReadString(t, filepath.Join(output, "sub", "sub", "5.txt"))
		assert.Equal(t, "", replaced)
	}
}

func TestReplaceCmd_Escape_String(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input.txt", "1\n2\n")
	output := filepath.Join(d, "output.txt")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-s", `\n`,
		"-t", `\t`,
		"--escape",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	replaced := test.ReadString(t, output)
	assert.Equal(t, "1\t2\t", replaced)
}

func TestReplaceCmd_Escape_Regex(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input.txt", "a　　　")
	output := filepath.Join(d, "output.txt")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-r", `\u3000+`,
		"-t", `\u0020`,
		"--escape",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	replaced := test.ReadString(t, output)
	assert.Equal(t, "a ", replaced)
}

func TestReplaceCmd_Encoding_UTF8(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input.txt", "あいうえお")
	output := filepath.Join(d, "output.txt")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-r", "あ.う",
		"-t", "",
		"--encoding", "utf-8",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	replaced := test.ReadString(t, output)
	assert.Equal(t, "えお", replaced)
}

func TestReplaceCmd_Encoding_SJIS(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteBytes(t, d, "input.txt", test.StringToByte(t, "あいうえお", japanese.ShiftJIS))
	output := filepath.Join(d, "output.txt")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-r", "あ.う",
		"-t", "",
		"--encoding", "sjis",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	replaced := test.ByteToString(t, test.ReadBytes(t, output), japanese.ShiftJIS)
	assert.Equal(t, "えお", replaced)
}

func TestReplaceCmd_Encoding_Binary(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteBytes(t, d, "input.txt", []byte{0x00, 0x01, 0x00, 0x02, 0x00, 0x01, 0xF0})
	output := filepath.Join(d, "output.txt")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-s", "x00x01",
		"-t", "",
		"--encoding", "binary",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.NoError(t, err)

	replaced := test.ReadBytes(t, output)
	assert.Equal(t, []byte{0x00, 0x02, 0xF0}, replaced)
}

func TestReplaceCmd_Encoding_Invalid(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input.txt", "")
	output := filepath.Join(d, "output.txt")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-r", "x",
		"-t", "",
		"--encoding", "xxxx", // 存在しないEncoding
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.Error(t, err)
	assert.EqualError(t, err, "xxxx is invalid: htmlindex: invalid encoding name")
}

func TestReplaceCmd_InvalidRegex(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input.txt", "")
	output := filepath.Join(d, "output.txt")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-r", "[a", // 不正な正規表現
		"-t", "",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.Error(t, err)
	assert.Equal(t, "regular expression specified in --regex is invalid: error parsing regexp: missing closing ]: `[a`", err.Error())
}

func TestReplaceCmd_InvalidEscape_Regex(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input.txt", "")
	output := filepath.Join(d, "output.txt")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-r", `\x`, // 不正なエスケープ
		"-t", "",
		"--escape",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.Error(t, err)
	assert.Equal(t, "could not parse value \\x of flag regex: invalid syntax", err.Error())
}

func TestReplaceCmd_InvalidEscape_String(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input.txt", "")
	output := filepath.Join(d, "output.txt")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-s", `\x`, // 不正なエスケープ
		"-t", "",
		"--escape",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.Error(t, err)
	assert.Equal(t, "could not parse value \\x of flag string: invalid syntax", err.Error())
}

func TestReplaceCmd_InvalidEscape_Replacement(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateFileWriteString(t, d, "input.txt", "")
	output := filepath.Join(d, "output.txt")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-s", "a",
		"-t", `\`, // 不正なエスケープ
		"--escape",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.Error(t, err)
	assert.Equal(t, "could not parse value \\ of flag replacement: invalid syntax", err.Error())
}

func TestReplaceCmd_InputNotFound(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := filepath.Join(d, "input") // 存在しない
	output := filepath.Join(d, "output")

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-s", "a",
		"-t", "",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.Error(t, err)
	// 実行環境によってファイルが存在しない場合のエラーメッセージが異なるので、Errorだけで判定
}

func TestReplaceCmd_OutputNotFound(t *testing.T) {

	// ARRANGE
	d := test.CreateTempDir(t)
	defer os.RemoveAll(d)

	input := test.CreateDir(t, d, "input")
	output := filepath.Join(d, "a", "b") // 親ディレクトリ自体が無い

	rootCmd := newRootCmd()
	rootCmd.SetArgs([]string{
		"replace",
		"-i", input,
		"-s", "a",
		"-t", "",
		"-o", output,
	})

	// ACT
	err := rootCmd.Execute()

	// ASSERT
	require.Error(t, err)
	// 実行環境によってファイルが存在しない場合のエラーメッセージが異なるので、Errorだけで判定
}
