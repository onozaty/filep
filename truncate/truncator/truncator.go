package truncator

import (
	"os"

	"github.com/onozaty/filep/extract/extractor"
)

type Truncator struct {
	extractor extractor.Extractor
}

func (t *Truncator) Truncate(inputFilePath string, outputFilePath string) error {
	return t.extractor.Extract(inputFilePath, outputFilePath)
}

// 空ファイルを作成するだけのExtractorです。
type emptyExtractor struct {
}

func (t *emptyExtractor) Extract(inputFilePath string, outputFilePath string) error {

	out, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	return out.Close()
}

func newEmptyTruncator() (*Truncator, error) {
	return &Truncator{
		extractor: &emptyExtractor{},
	}, nil
}
