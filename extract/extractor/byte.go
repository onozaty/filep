package extractor

import (
	"fmt"
	"io"
	"os"
)

type byteExtractor struct {
	start int64
	end   int64
}

func NewByteExtractor(start int64, end int64) (Extractor, error) {

	if start < 1 || end < start {
		return nil, fmt.Errorf("invalid range: start = %d, end = %d", start, end)
	}

	return &byteExtractor{
		start: start,
		end:   end,
	}, nil
}

func (t *byteExtractor) Extract(inputFilePath string, outputFilePath string) error {

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

	// 開始位置を変更
	input.Seek(t.start-1, 0)
	_, err = io.CopyN(out, input, t.end-t.start+1)
	if err != nil && err != io.EOF { // 入力ファイルが指定サイズ未満の場合はEOFが返される(そこまでの書き込みでOK)
		return err
	}

	return nil
}
