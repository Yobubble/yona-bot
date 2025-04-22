package storage

import (
	"os"
	"path/filepath"
)

type localStorage struct {
}

func (l *localStorage) Read(filePath string) ([]byte, error) {
	fileByte, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return fileByte, nil
}

func (l *localStorage) Write(data []byte, filePath string) error {
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

func newLocal() Storage {
	return &localStorage{}
}
