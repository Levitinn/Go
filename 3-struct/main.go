package main

import (
	"3-struct/api"
	"3-struct/bins"
	"3-struct/config"
	"3-struct/file"
	"3-struct/storage"
	"fmt"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	client := api.NewClient(cfg)
	fmt.Printf("API client ready, key loaded (%d chars)\n", len(client.APIKey()))

	newFile := file.NewFile("storage.json")
	newStorage := storage.NewStorage(newFile)
	binsList := bins.NewBins()
	bin := binsList.NewBin("test", false)

	err = newStorage.Save(bin)
	if err != nil {
		fmt.Println("Error saving bin:", err)
		return
	}
	list, err := newStorage.Load()
	if err != nil {
		fmt.Println("Error loading bins:", err)
		return
	}
	fmt.Println(list)

}
