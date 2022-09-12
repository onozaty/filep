package encoding

import (
	"github.com/pkg/errors"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/htmlindex"
)

func Encoding(name string) (encoding.Encoding, error) {

	encoding, err := htmlindex.Get(name)
	if err != nil {
		return nil, errors.WithMessagef(err, "%s is invalid", name)
	}

	return encoding, nil
}
