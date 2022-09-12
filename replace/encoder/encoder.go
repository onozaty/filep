package encoder

import (
	"strings"

	enc "github.com/onozaty/filep/encoding"

	"golang.org/x/text/encoding"
)

type Encoder interface {
	String([]byte) (string, error)
	Bytes(string) ([]byte, error)
}

func NewEncoder(name string) (Encoder, error) {

	if strings.ToLower(name) == "binary" {
		return &BinaryEncoder{}, nil
	}

	return newEncodingEncoder(name)
}

func newEncodingEncoder(name string) (*EncodingEncoder, error) {

	encoding, err := enc.Encoding(name)
	if err != nil {
		return nil, err
	}

	return &EncodingEncoder{
		encoding: encoding,
	}, nil
}

type EncodingEncoder struct {
	encoding encoding.Encoding
}

func (e *EncodingEncoder) String(src []byte) (string, error) {

	decodedBytes, err := e.encoding.NewDecoder().Bytes(src)
	if err != nil {
		return "", err
	}

	return string(decodedBytes), nil
}

func (e *EncodingEncoder) Bytes(src string) ([]byte, error) {

	return e.encoding.NewEncoder().Bytes([]byte(src))
}
