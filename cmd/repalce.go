package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/onozaty/filep/replace/encoder"
	"github.com/onozaty/filep/replace/replacer"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func newReplaceCmd() *cobra.Command {

	replaceCmd := &cobra.Command{
		Use:   "replace",
		Short: "Replace file contents",
		RunE: func(cmd *cobra.Command, args []string) error {

			inputPath, _ := cmd.Flags().GetString("input")
			outputPath, _ := cmd.Flags().GetString("output")

			escapeSequence, _ := cmd.Flags().GetBool("escape")

			// エスケープ対象のフラグはエスケープ有無に応じて変換
			targetRegex, err := getFlagEscapedString(cmd.Flags(), "regex", escapeSequence)
			if err != nil {
				return err
			}

			targetStr, err := getFlagEscapedString(cmd.Flags(), "string", escapeSequence)
			if err != nil {
				return err
			}

			replacement, err := getFlagEscapedString(cmd.Flags(), "replacement", escapeSequence)
			if err != nil {
				return err
			}

			recursive, _ := cmd.Flags().GetBool("recursive")
			encoding, _ := cmd.Flags().GetString("encoding")

			if targetStr == "" && targetRegex == "" {
				return fmt.Errorf("--regex or --string must be specified")
			}

			var regex *regexp.Regexp
			if targetRegex != "" {
				regex, err = regexp.Compile(targetRegex)
				if err != nil {
					return errors.WithMessage(err, "regular expression specified in --regex is invalid")
				}
			}

			// 引数の解析に成功した時点で、エラーが起きてもUsageは表示しない
			cmd.SilenceUsage = true

			return runReplace(
				inputPath,
				outputPath,
				replaceCondition{
					targetRegex: regex,
					targetStr:   targetStr,
					replacement: replacement,
				},
				encoding,
				recursive)
		},
	}

	replaceCmd.Flags().StringP("input", "i", "", "Input file/dir path.")
	replaceCmd.MarkFlagRequired("input")
	replaceCmd.Flags().StringP("output", "o", "", "Output file/dir path.")
	replaceCmd.MarkFlagRequired("output")

	replaceCmd.Flags().StringP("regex", "r", "", "Target regex.")
	replaceCmd.Flags().StringP("string", "s", "", "Target string.")
	replaceCmd.Flags().StringP("replacement", "t", "", "Replacement.")
	replaceCmd.MarkFlagRequired("replacement")

	replaceCmd.Flags().BoolP("escape", "", false, "Enable escape sequence.")
	replaceCmd.Flags().BoolP("recursive", "", false, "Recursively traverse the input dir.")
	replaceCmd.Flags().StringP("encoding", "", "UTF-8", "Encoding.")

	return replaceCmd
}

type replaceCondition struct {
	targetRegex *regexp.Regexp
	targetStr   string
	replacement string
}

func runReplace(inputPath string, outputPath string, condition replaceCondition, encoding string, recursive bool) error {

	encoder, err := encoder.NewEncoder(encoding)
	if err != nil {
		return err
	}

	replacer := newReplacer(condition)

	process := func(inputFilePath string, outputFilePath string) error {
		return replaceFile(inputFilePath, outputFilePath, replacer, encoder)
	}

	return handle(inputPath, outputPath, process, recursive)
}

func replaceFile(inputFilePath string, outputFilePath string, replacer replacer.Replacer, encoder encoder.Encoder) error {

	inputBytes, err := os.ReadFile(inputFilePath)
	if err != nil {
		return err
	}

	inputContents, err := encoder.String(inputBytes)
	if err != nil {
		return err
	}

	outputContents := replacer.Replace(inputContents)

	out, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer out.Close()

	encodedBytes, err := encoder.Bytes(outputContents)
	if err != nil {
		return err
	}

	_, err = out.Write(encodedBytes)
	return err
}

func newReplacer(condition replaceCondition) replacer.Replacer {

	if condition.targetRegex != nil {
		return replacer.NewRegexpReplacer(condition.targetRegex, condition.replacement)
	}

	return replacer.NewStringReplacer(condition.targetStr, condition.replacement)
}

func getFlagEscapedString(f *pflag.FlagSet, name string, escape bool) (string, error) {

	str, _ := f.GetString(name)

	if !escape {
		// エスケープ無しの場合
		return str, nil
	}

	// \nのように指定されているものを、スケープ文字として扱えるように
	unq, err := strconv.Unquote(`"` + str + `"`)
	if err != nil {
		return "", errors.Wrapf(err, "could not parse value %s of flag %s", str, name)
	}

	return unq, nil
}
