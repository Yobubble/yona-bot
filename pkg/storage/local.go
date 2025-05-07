package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

type localStorage struct {
}

func (l *localStorage) Read(filePath string) ([]byte, error) {
	fileByte, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("storage: error reading file %s: %w", filePath, err)
	}
	return fileByte, nil
}

func (l *localStorage) Write(data []byte, filePath string) error {
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return fmt.Errorf("storage: error creating directory %s: %w", filepath.Dir(filePath), err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("storage: error creating file %s: %w", filePath, err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("storage: error writing to file %s: %w", filePath, err)
	}

	return nil
}

func newLocal() Storage {
	return &localStorage{}
}
