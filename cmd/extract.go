package cmd

import (
	"fmt"
	"math"

	"github.com/onozaty/filep/extract/extractor"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func newExtractCmd() *cobra.Command {

	extractCmd := &cobra.Command{
		Use:   "extract",
		Short: "Extract file contents",
		RunE: func(cmd *cobra.Command, args []string) error {

			inputPath, _ := cmd.Flags().GetString("input")
			outputPath, _ := cmd.Flags().GetString("output")

			start := getFlagInt64(cmd.Flags(), "start", 1)
			// endの指定が無かった場合には、ファイル終端までを対象にするためにint64の最大値を入れておく
			end := getFlagInt64(cmd.Flags(), "end", math.MaxInt64)

			countingType, err := getFlagCountingType(cmd.Flags())
			if err != nil {
				return err
			}

			recursive, _ := cmd.Flags().GetBool("recursive")
			encoding, _ := cmd.Flags().GetString("encoding")

			if start <= 0 {
				return fmt.Errorf("start must be greater than or equal to 1")
			}
			if start > end {
				return fmt.Errorf("end must be greater than or equal to start")
			}

			// 引数の解析に成功した時点で、エラーが起きてもUsageは表示しない
			cmd.SilenceUsage = true

			return runExtract(
				inputPath,
				outputPath,
				extractCondition{
					start:        start,
					end:          end,
					countingType: countingType,
				},
				encoding,
				recursive)
		},
	}

	extractCmd.Flags().StringP("input", "i", "", "Input file/dir path.")
	extractCmd.MarkFlagRequired("input")
	extractCmd.Flags().StringP("output", "o", "", "Output file/dir path.")
	extractCmd.MarkFlagRequired("output")

	extractCmd.Flags().Int64P("start", "s", 0, "Start position.")
	extractCmd.Flags().Int64P("end", "e", 0, "End position.")

	extractCmd.Flags().BoolP("byte", "b", false, "Handle by bytes.")
	extractCmd.Flags().BoolP("char", "c", false, "Handle by characters.")
	extractCmd.Flags().BoolP("line", "l", false, "Handle by lines.")

	extractCmd.Flags().BoolP("recursive", "", false, "Recursively traverse the input dir.")
	extractCmd.Flags().StringP("encoding", "", "UTF-8", "Encoding.")

	return extractCmd
}

type extractCondition struct {
	start        int64
	end          int64
	countingType CountingType
}

func runExtract(inputPath string, outputPath string, condition extractCondition, encoding string, recursive bool) error {

	extractor, err := newExtractor(condition, encoding)
	if err != nil {
		return err
	}

	process := func(inputFilePath string, outputFilePath string) error {
		return extractor.Extract(inputFilePath, outputFilePath)
	}

	return handle(inputPath, outputPath, process, recursive)
}

func newExtractor(condition extractCondition, encoding string) (extractor.Extractor, error) {

	switch condition.countingType {
	case Bytes:
		return extractor.NewByteExtractor(condition.start, condition.end)
	case Chars:
		return extractor.NewCharExtractor(condition.start, condition.end, encoding)
	case Lines:
		return extractor.NewLineExtractor(condition.start, condition.end, encoding)
	default:
		return nil, fmt.Errorf("invalid counting type: %d", condition.countingType)
	}
}

func getFlagInt64(f *pflag.FlagSet, name string, defaultValue int64) int64 {

	if f.Changed(name) {
		val, _ := f.GetInt64(name)
		return val
	}

	return defaultValue
}
