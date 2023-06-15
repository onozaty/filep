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

type charExtractor struct {
	start    int64
	end      int64
	encoding encoding.Encoding
}

func NewCharExtractor(start int64, end int64, encodingName string) (Extractor, error) {

	if start < 1 || end < start {
		return nil, fmt.Errorf("invalid range: start = %d, end = %d", start, end)
	}

	encoding, err := enc.Encoding(encodingName)
	if err != nil {
		return nil, err
	}

	return &charExtractor{
		start:    start,
		end:      end,
		encoding: encoding,
	}, nil
}

func (t *charExtractor) Extract(inputFilePath string, outputFilePath string) error {

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

	// 指定範囲を取り出し
	for currentCharNum := int64(1); currentCharNum <= t.end; currentCharNum++ {
		c, _, err := reader.ReadRune()
		if err == io.EOF {
			// 終端ならばそこまでで終了
			break
		}
		if err != nil {
			return err
		}

		if currentCharNum >= t.start {
			// 開始位置を満たしていたら出力
			if _, err := writer.WriteRune(c); err != nil {
				return err
			}
		}
	}

	return writer.Flush()
}
