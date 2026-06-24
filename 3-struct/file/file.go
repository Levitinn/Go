package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func IsJSON(fileName string) bool {
	fileExt := filepath.Ext(strings.ToLower(fileName))
	return fileExt == ".json"
}

// Функция для записи в файл
func WriteFile(content []byte, name string) error {
	file, err := os.Create(name)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	} else {
		fmt.Println("File written successfully")
	}
	return nil
}
