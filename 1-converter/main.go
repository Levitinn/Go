package main

import "fmt"

const (
	USDToEUR = 0.85
	USDToRUB = 75.0
)

const EURToRUB = USDToRUB / USDToEUR

func main() {
	amount, fromCurrency, toCurrency := readInput()
	result := calcResult(amount, fromCurrency, toCurrency)
	fmt.Printf("%.2f\n", result)
}

func readInput() (float64, string, string) {
	var amount float64
	var fromCurrency, toCurrency string

	fmt.Print("Enter amount: ")
	fmt.Scan(&amount)

	fmt.Print("Enter source currency (USD/EUR/RUB): ")
	fmt.Scan(&fromCurrency)

	fmt.Print("Enter target currency (USD/EUR/RUB): ")
	fmt.Scan(&toCurrency)

	return amount, fromCurrency, toCurrency
}

func calcResult(amount float64, fromCurrency, toCurrency string) float64 {
	return 0
}
