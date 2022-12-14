package cmd

import (
	"fmt"

	"github.com/onozaty/filep/truncate/truncator"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func newTruncateCmd() *cobra.Command {

	truncateCmd := &cobra.Command{
		Use:   "truncate",
		Short: "Truncate file contents",
		RunE: func(cmd *cobra.Command, args []string) error {

			inputPath, _ := cmd.Flags().GetString("input")
			outputPath, _ := cmd.Flags().GetString("output")

			byteNum := getFlagTruncateNum(cmd.Flags(), "byte")
			charNum := getFlagTruncateNum(cmd.Flags(), "char")
			lineNum := getFlagTruncateNum(cmd.Flags(), "line")

			recursive, _ := cmd.Flags().GetBool("recursive")
			encoding, _ := cmd.Flags().GetString("encoding")

			if byteNum == nil && charNum == nil && lineNum == nil {
				return fmt.Errorf("no number is specified")
			}

			// 引数の解析に成功した時点で、エラーが起きてもUsageは表示しない
			cmd.SilenceUsage = true

			return runTruncate(
				inputPath,
				outputPath,
				truncateCondition{
					byteNum: byteNum,
					charNum: charNum,
					lineNum: lineNum,
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
	byteNum *int64
	charNum *int64
	lineNum *int64
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

func newTruncator(condition truncateCondition, encoding string) (truncator.Truncator, error) {

	if condition.byteNum != nil {
		return truncator.NewByteTruncator(*condition.byteNum), nil
	} else if condition.charNum != nil {
		return truncator.NewCharTruncator(*condition.charNum, encoding)
	} else {
		return truncator.NewLineTruncator(*condition.lineNum, encoding)
	}
}

func getFlagTruncateNum(f *pflag.FlagSet, name string) *int64 {

	if f.Changed(name) {
		num, _ := f.GetInt64(name)
		return &num
	}

	return nil
}
