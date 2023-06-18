package cmd

import (
	"fmt"

	"github.com/spf13/pflag"
)

type CountingType int

const (
	Bytes CountingType = iota
	Chars
	Lines
)

func getFlagCountingType(f *pflag.FlagSet) (CountingType, error) {

	handleByte, _ := f.GetBool("byte")
	handleChar, _ := f.GetBool("char")
	handleLine, _ := f.GetBool("line")

	selected := 0
	if handleByte {
		selected++
	}
	if handleChar {
		selected++
	}
	if handleLine {
		selected++
	}

	if selected != 1 {
		return 0, fmt.Errorf("specify one of the following: -b, -c, -l")
	}

	if handleByte {
		return Bytes, nil
	}
	if handleChar {
		return Chars, nil
	}
	return Lines, nil
}

func getFlagCountingTypeWithNumber(f *pflag.FlagSet) (CountingType, int64, error) {

	byteNum := getFlagNum(f, "byte")
	charNum := getFlagNum(f, "char")
	lineNum := getFlagNum(f, "line")

	var selectNum int64
	var countingType CountingType

	selected := 0
	if byteNum != nil {
		selected++
		selectNum = *byteNum
		countingType = Bytes
	}
	if charNum != nil {
		selected++
		selectNum = *charNum
		countingType = Chars
	}
	if lineNum != nil {
		selected++
		selectNum = *lineNum
		countingType = Lines
	}

	if selected != 1 {
		return 0, 0, fmt.Errorf("specify one of the following: -b, -c, -l")
	}

	return countingType, selectNum, nil
}

func getFlagNum(f *pflag.FlagSet, name string) *int64 {

	if f.Changed(name) {
		num, _ := f.GetInt64(name)
		return &num
	}

	return nil
}
