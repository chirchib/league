package logger

import (
	"errors"
	"os"
	"path/filepath"
)

type fileWriter struct {
	file     *os.File
	filePath string
}

func NewFileWriter(filePath string) (*fileWriter, error) {
	file, err := getOrCreateFileIfNotExists(filePath)
	if err != nil {
		return nil, err
	}

	return &fileWriter{
		file:     file,
		filePath: filePath,
	}, nil
}

func (fw *fileWriter) Write(b []byte) (int, error) {
	var err error

	if _, err = os.Stat(fw.filePath); err != nil {
		fw.file, err = getOrCreateFileIfNotExists(fw.filePath)
		if err != nil {
			return 0, err
		}
	}

	return fw.file.Write(b)
}

func getOrCreateFileIfNotExists(filePath string) (*os.File, error) {
	var (
		file *os.File
		err  error
	)

	dirPath := filepath.Dir(filePath)

	if _, err = os.Stat(dirPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}
