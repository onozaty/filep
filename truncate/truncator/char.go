package truncator

import "github.com/onozaty/filep/extract/extractor"

func NewCharTruncator(charNum int64, encodingName string) (*Truncator, error) {

	if charNum == 0 {
		// 0を指定された場合、空ファイルを作るだけ
		return newEmptyTruncator()
	}

	// 1文字目から取り出すことで切り捨てと同じ扱いに
	extractor, err := extractor.NewCharExtractor(1, charNum, encodingName)
	if err != nil {
		return nil, err
	}

	return &Truncator{
		extractor: extractor,
	}, nil
}
