package replacer

import (
	"regexp"
)

type regexpReplacer struct {
	regex       *regexp.Regexp
	replacement string
}

func NewRegexpReplacer(regex *regexp.Regexp, replacement string) Replacer {

	return &regexpReplacer{
		regex:       regex,
		replacement: replacement,
	}
}

func (r *regexpReplacer) Replace(s string) string {
	return r.regex.ReplaceAllString(s, r.replacement)
}
