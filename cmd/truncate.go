package cmd

import (
	"fmt"

	"github.com/onozaty/filep/truncate/truncator"

	"github.com/spf13/cobra"
)

func newTruncateCmd() *cobra.Command {

	truncateCmd := &cobra.Command{
		Use:   "truncate",
		Short: "Truncate file contents",
		RunE: func(cmd *cobra.Command, args []string) error {

			inputPath, _ := cmd.Flags().GetString("input")
			outputPath, _ := cmd.Flags().GetString("output")

			countingType, number, err := getFlagCountingTypeWithNumber(cmd.Flags())
			if err != nil {
				return err
			}

			recursive, _ := cmd.Flags().GetBool("recursive")
			encoding, _ := cmd.Flags().GetString("encoding")

			if number < 0 {
				return fmt.Errorf("number must be greater than or equal to 0")
			}

			// 引数の解析に成功した時点で、エラーが起きてもUsageは表示しない
			cmd.SilenceUsage = true

			return runTruncate(
				inputPath,
				outputPath,
				truncateCondition{
					countingType: countingType,
					number:       number,
				},
				encoding,
				recursive)
		},
	}

	truncateCmd.Flags().StringP("input", "i", "", "Input file/dir path.")
	truncateCmd.MarkFlagRequired("input")
	truncateCmd.Flags().StringP("output", "o", "", "Output file/dir path.")
	truncateCmd.MarkFlagRequired("output")

	truncateCmd.Flags().Int64P("byte", "b", 0, "Number of bytes.")
	truncateCmd.Flags().Int64P("char", "c", 0, "Number of characters.")
	truncateCmd.Flags().Int64P("line", "l", 0, "Number of lines.")

	truncateCmd.Flags().BoolP("recursive", "", false, "Recursively traverse the input dir.")
	truncateCmd.Flags().StringP("encoding", "", "UTF-8", "Encoding.")

	return truncateCmd
}

type truncateCondition struct {
	countingType CountingType
	number       int64
}

func runTruncate(inputPath string, outputPath string, condition truncateCondition, encoding string, recursive bool) error {

	truncator, err := newTruncator(condition, encoding)
	if err != nil {
		return err
	}

	process := func(inputFilePath string, outputFilePath string) error {
		return truncator.Truncate(inputFilePath, outputFilePath)
	}

	return handle(inputPath, outputPath, process, recursive)
}

func newTruncator(condition truncateCondition, encoding string) (*truncator.Truncator, error) {

	switch condition.countingType {
	case Bytes:
		return truncator.NewByteTruncator(condition.number)
	case Chars:
		return truncator.NewCharTruncator(condition.number, encoding)
	case Lines:
		return truncator.NewLineTruncator(condition.number, encoding)
	default:
		return nil, fmt.Errorf("invalid counting type: %d", condition.countingType)
	}
}
