package replacer

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexpReplacer(t *testing.T) {

	replacer := NewRegexpReplacer(regexp.MustCompile("[a-z]{2}"), "xx")

	{
		result := replacer.Replace("abc")
		assert.Equal(t, "xxc", result)
	}
	{
		result := replacer.Replace("abcd")
		assert.Equal(t, "xxxx", result)
	}
	{
		result := replacer.Replace("a")
		assert.Equal(t, "a", result)
	}
	{
		result := replacer.Replace("")
		assert.Equal(t, "", result)
	}
}

func TestRegexpReplacer_BackReference(t *testing.T) {

	replacer := NewRegexpReplacer(regexp.MustCompile("X([0-9]+)"), "Z$1")

	{
		result := replacer.Replace("X123X")
		assert.Equal(t, "Z123X", result)
	}
	{
		result := replacer.Replace("X1X2X3X")
		assert.Equal(t, "Z1Z2Z3X", result)
	}
	{
		result := replacer.Replace("a")
		assert.Equal(t, "a", result)
	}
	{
		result := replacer.Replace("")
		assert.Equal(t, "", result)
	}
}
