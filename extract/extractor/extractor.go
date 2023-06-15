package extractor

type Extractor interface {
	Extract(inputFilePath string, outputFilePath string) error
}
