package truncator

import (
	"bufio"
	"io"
	"os"

	enc "github.com/onozaty/filep/encoding"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

type charTruncator struct {
	charNum  int64
	encoding encoding.Encoding
}

func NewCharTruncator(charNum int64, charset string) (Truncator, error) {

	encoding, err := enc.Encoding(charset)
	if err != nil {
		return nil, err
	}

	return &charTruncator{
		charNum:  charNum,
		encoding: encoding,
	}, nil
}

func (t *charTruncator) Truncate(inputFilePath string, outputFilePath string) error {

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

	// 指定文字数分読み込み
	for i := int64(0); i < t.charNum; i++ {
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
	}

	return writer.Flush()
}
