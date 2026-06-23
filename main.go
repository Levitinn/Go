package main

import (
	"fmt"
	"strings"
)

func main() {
	string := "1, 10, 23, 41, 112,12,2"
	parts := strings.Split(string, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
	}
	fmt.Println(parts)
}
