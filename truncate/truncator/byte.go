package truncator

import (
	"io"
	"os"
)

type byteTruncator struct {
	byteNum int64
}

func NewByteTruncator(byteNum int64) Truncator {

	return &byteTruncator{
		byteNum: byteNum,
	}
}

func (t *byteTruncator) Truncate(inputFilePath string, outputFilePath string) error {

	input, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}
	defer input.Close()

	out, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 指定サイズ分コピー
	_, err = io.CopyN(out, input, t.byteNum)
	if err != nil && err != io.EOF { // 入力ファイルが指定サイズ未満の場合はEOFが返される(そこまでの書き込みでOK)
		return err
	}

	return nil
}
