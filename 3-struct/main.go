package main

import "time"

type BinList []Bin

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func newBin(name string, private bool) *Bin {
	return &Bin{
		name:    name,
		private: private,
	}
}

func main() {

}
