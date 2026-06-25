package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type File interface {
	Read() ([]byte, error)
	Write([]byte) error
	IsJSON() bool
}
type fileImpl struct {
	path string
}

func NewFile(path string) File {
	return &fileImpl{path: path}
}
func (file *fileImpl) Read() ([]byte, error) {
	data, err := os.ReadFile(file.path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (file *fileImpl) IsJSON() bool {
	fileExt := filepath.Ext(strings.ToLower(file.path))
	return fileExt == ".json"
}

// Функция для записи в файл
func (f *fileImpl) Write(data []byte) error {
	file, err := os.Create(f.path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return nil
}
