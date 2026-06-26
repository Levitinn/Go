package main

import "fmt"

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}
type jsonDb struct {
	data []byte
}

func (db *jsonDb) Read() ([]byte, error) {
	return db.data, nil
}
func (db *jsonDb) Write(data []byte) {
	db.data = make([]byte, len(data))
	copy(db.data, data)
}
func main() {
	db := &jsonDb{}
	db.Write([]byte("Hello, World!"))
	data, err := db.Read()
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}
	fmt.Println(string(data))
}
