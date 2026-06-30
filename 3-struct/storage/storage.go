package storage

import (
	"3-struct/bins"
	"3-struct/file"
	"encoding/json"
	"fmt"
	"os"
)

const FileName = "storage.json"

type Storage interface {
	Save(bin *bins.Bin) error
	Load() (bins.BinList, error)
}

type jsonStorage struct {
	file file.File
}

func NewStorage(file file.File) Storage {
	return &jsonStorage{file: file}
}

func (storage *jsonStorage) Save(bin *bins.Bin) error {
	list := bins.BinList{}
	if !storage.file.IsJSON() {
		return fmt.Errorf("file %s is not a JSON", FileName)
	}
	data, err := storage.file.Read()
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else if len(data) > 0 {
		err = json.Unmarshal(data, &list)
		if err != nil {
			return err
		}
	}
	list = append(list, *bin)
	bytes, err := json.Marshal(list)
	if err != nil {
		return err
	}
	err = storage.file.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func (storage *jsonStorage) Load() (bins.BinList, error) {
	binList := bins.BinList{}
	if !storage.file.IsJSON() {
		return binList, fmt.Errorf("file %s is not a JSON")
	}
	content, err := storage.file.Read()
	if err != nil {
		if !os.IsNotExist(err) {
			return binList, err
		}
		return binList, nil
	}
	err = json.Unmarshal(content, &binList)
	if err != nil {
		return nil, err
	}
	return binList, nil

}
