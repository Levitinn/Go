package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randChan := make(chan int)
	squareChan := make(chan int)
	go createRandomSlice(randChan)
	go squareNumbers(randChan, squareChan)
	for num := range squareChan {
		fmt.Println(num)
	}
}

func createRandomSlice(ch chan int) {
	for i := 0; i < 10; i++ {
		randomInt := rand.Intn(100)
		ch <- randomInt
	}
	close(ch)
}
func squareNumbers(in chan int, out chan int) {
	for number := range in {
		out <- number * number
	}
	close(out)
}
