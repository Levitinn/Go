package main

import "fmt"

type stringMap map[string]string

func main() {
	fruitMap := make(stringMap, 5)
	fruitMap["apple"] = "яблоко"
	fruitMap["banana"] = "банан"
	fruitMap["cherry"] = "вишня"
	startLength := len(fruitMap)
	lastKey := "grape"
	fruitMap[lastKey] = "виноград"
	fmt.Printf("len before=%d after=%d\n", startLength, len(fruitMap))
	fmt.Printf(fruitMap[string(len(fruitMap)-1)])
	fmt.Printf("%s: %s \n", lastKey, fruitMap[lastKey])
}
