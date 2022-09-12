package truncator

type Truncator interface {
	Truncate(inputFilePath string, outputFilePath string) error
}
