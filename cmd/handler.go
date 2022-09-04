package cmd

import (
	"os"
	"path/filepath"
)

func handle(inputPath string, outputPath string, process func(inputFilePath string, outputFilePath string) error, recursive bool) error {

	inputInfo, err := os.Stat(inputPath)
	if err != nil {
		return err
	}

	if !inputInfo.IsDir() {
		// ファイル指定
		return handleFile(inputPath, outputPath, process)
	} else {
		// ディレクトリ指定
		return handleFiles(inputPath, outputPath, process, recursive)
	}
}

func handleFiles(inputDirPath string, outputDirPath string, process func(inputFilePath string, outputFilePath string) error, recursive bool) error {

	entries, err := os.ReadDir(inputDirPath)
	if err != nil {
		return err
	}

	// 出力先のディレクトリが無かったら作っておく
	_, err = os.Stat(outputDirPath)
	if os.IsNotExist(err) {
		if err := os.Mkdir(outputDirPath, os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			err := handleFile(filepath.Join(inputDirPath, entry.Name()), filepath.Join(outputDirPath, entry.Name()), process)
			if err != nil {
				return err
			}
		} else if recursive {
			// ディレクトリかつ再帰的にたどる場合
			if err := handleFiles(filepath.Join(inputDirPath, entry.Name()), filepath.Join(outputDirPath, entry.Name()), process, recursive); err != nil {
				return err
			}
		}
	}

	return nil
}

func handleFile(inputFilePath string, outputFilePath string, process func(inputPath string, outputPath string) error) error {

	return process(inputFilePath, outputFilePath)
}
