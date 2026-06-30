package main

import (
	"3-struct/api"
	"3-struct/bins"
	"3-struct/config"
	"3-struct/file"
	"3-struct/storage"
	"flag"
	"fmt"
	"os"
)

var create = flag.Bool("create", false, "create a new bin")
var update = flag.Bool("update", false, "update a bin")
var delete = flag.Bool("delete", false, "delete a bin")
var get = flag.Bool("get", false, "get a bin")
var list = flag.Bool("list", false, "list all bins")

var filePath = flag.String("file", "", "path to the storage file")
var name = flag.String("name", "", "name of the bin")
var id = flag.String("id", "", "id of the bin")

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	flag.Parse()
	client := api.NewClient(cfg)
	store := storage.NewStorage(file.NewFile("storage.json"))
	switch {
	case *list:
		list, err := store.Load()
		if err != nil {
			fmt.Println("Error loading bins:", err)
			return
		}
		for _, bin := range list {
			fmt.Println(bin.ID, bin.Name)
		}
	case *get:
		record, err := client.GetBin(*id)
		if err != nil {
			fmt.Println("Error getting bin:", err)
			return
		}
		fmt.Println(string(record))
	case *delete:
		err := client.DeleteBin(*id)
		if err != nil {
			fmt.Println("Error deleting bin:", err)
			return
		}
		err = store.Delete(*id)
		if err != nil {
			fmt.Println("Error deleting bin:", err)
			return
		}
		fmt.Println("Bin deleted successfully")
	case *create:
		data, err := os.ReadFile(*filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		id, err := client.CreateBin(data, *name)
		if err != nil {
			fmt.Println("Error creating bin:", err)
			return
		}
		store.Save(bins.Bin{ID: id, Name: *name})
		fmt.Println("Bin created successfully with id:", id)
	case *update:
		data, err := os.ReadFile(*filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		err = client.UpdateBin(*id, data)
		if err != nil {
			fmt.Println("Error updating bin:", err)
			return
		}
		fmt.Println("Bin updated successfully")
	default:
		flag.Usage()
	}

}
