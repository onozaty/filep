package truncator

import (
	"bufio"
	"io"
	"os"

	enc "github.com/onozaty/filep/encoding"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

type lineTruncator struct {
	lineNum  int64
	encoding encoding.Encoding
}

func NewLineTruncator(lineNum int64, encodingName string) (Truncator, error) {

	encoding, err := enc.Encoding(encodingName)
	if err != nil {
		return nil, err
	}

	return &lineTruncator{
		lineNum:  lineNum,
		encoding: encoding,
	}, nil
}

func (t *lineTruncator) Truncate(inputFilePath string, outputFilePath string) error {

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

	// 指定行数分読み込み
	count := int64(0)
	for count < t.lineNum {
		c, _, err := reader.ReadRune()
		if err == io.EOF {
			// 終端ならばそこまでで終了
			break
		}
		if err != nil {
			return err
		}

		if _, err := writer.WriteRune(c); err != nil {
			return err
		}

		if c == '\n' {
			// LFで行をカウント
			count++
		}
	}

	return writer.Flush()
}
