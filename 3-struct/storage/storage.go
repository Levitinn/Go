package storage

import (
	"3-struct/bins"
	"3-struct/file"
	"encoding/json"
	"fmt"
	"os"
)

const FileName = "storage.json"

func SaveToStorage(bin *bins.Bin) error {
	list := bins.BinList{}
	if !file.IsJSON(FileName) {
		return fmt.Errorf("file %s is not a JSON", FileName)
	}
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

func LoadBins() (bins.BinList, error) {
	binList := bins.BinList{}
	if !file.IsJSON(FileName) {
		return binList, fmt.Errorf("file %s is not a JSON", FileName)
	}
	content, err := file.ReadFile(FileName)
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
