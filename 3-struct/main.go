package main

import (
	"3-struct/bins"
	"3-struct/file"
	"3-struct/storage"
)

func main() {
	file.ReadFile(storage.FileName)
	storage.SaveToStorage(bins.NewBin("test", false))

}
