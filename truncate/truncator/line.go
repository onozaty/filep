package truncator

import (
	"github.com/onozaty/filep/extract/extractor"
)

func NewLineTruncator(lineNum int64, encodingName string) (*Truncator, error) {

	if lineNum == 0 {
		// 0を指定された場合、空ファイルを作るだけ
		return newEmptyTruncator()
	}

	// 1行目から取り出すことで切り捨てと同じ扱いに
	extractor, err := extractor.NewLineExtractor(1, lineNum, encodingName)
	if err != nil {
		return nil, err
	}

	return &Truncator{
		extractor: extractor,
	}, nil
}
