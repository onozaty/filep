package truncator

import (
	"github.com/onozaty/filep/extract/extractor"
)

func NewByteTruncator(byteNum int64) (*Truncator, error) {

	if byteNum == 0 {
		// 0を指定された場合、空ファイルを作るだけ
		return newEmptyTruncator()
	}

	// 1バイト目から取り出すことで切り捨てと同じ扱いに
	extractor, err := extractor.NewByteExtractor(1, byteNum)
	if err != nil {
		return nil, err
	}

	return &Truncator{
		extractor: extractor,
	}, nil
}
