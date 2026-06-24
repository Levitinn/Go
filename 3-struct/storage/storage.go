package storage

import (
	"3-struct/bins"
	"3-struct/file"
	"encoding/json"
	"os"
)

const FileName = "storage.json"

func SaveToStorage(bin *bins.Bin) error {
	list := bins.BinList{}
	data, err := file.ReadFile(FileName)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
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
	err = file.WriteFile(bytes, FileName)
	if err != nil {
		return err
	}
	return nil
}
