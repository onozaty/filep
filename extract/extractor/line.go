package extractor

import (
	"bufio"
	"fmt"
	"io"
	"os"

	enc "github.com/onozaty/filep/encoding"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

type lineExtractor struct {
	start    int64
	end      int64
	encoding encoding.Encoding
}

func NewLineExtractor(start int64, end int64, encodingName string) (Extractor, error) {

	if start < 1 || end < start {
		return nil, fmt.Errorf("invalid range: start = %d, end = %d", start, end)
	}

	encoding, err := enc.Encoding(encodingName)
	if err != nil {
		return nil, err
	}

	return &lineExtractor{
		start:    start,
		end:      end,
		encoding: encoding,
	}, nil
}

func (t *lineExtractor) Extract(inputFilePath string, outputFilePath string) error {

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

	reader := bufio.NewReader(transform.NewReader(input, t.encoding.NewDecoder()))
	writer := bufio.NewWriter(transform.NewWriter(out, t.encoding.NewEncoder()))

	currentLineNum := int64(1) // 現在行は1行目から
	for currentLineNum <= t.end {
		c, _, err := reader.ReadRune()
		if err == io.EOF {
			// 終端ならばそこまでで終了
			break
		}
		if err != nil {
			return err
		}

		if currentLineNum >= t.start {
			// 開始位置を満たしていたら出力
			if _, err := writer.WriteRune(c); err != nil {
				return err
			}
		}

		if c == '\n' {
			// LFで行数をインクリメント
			currentLineNum++
		}
	}

	return writer.Flush()
}
