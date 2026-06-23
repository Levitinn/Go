package main

import "fmt"

const (
	USDToEUR = 0.85
	USDToRUB = 75.0
	EURToUSD = 1.17
)

const EURToRUB = USDToRUB / USDToEUR

func main() {
	amount, fromCurrency, toCurrency := readInput()
	result := calcResult(amount, fromCurrency, toCurrency)
	fmt.Printf("%.2f\n", result)
}

func isValidCurrency(currency string) bool {
	return currency == "USD" || currency == "EUR" || currency == "RUB"
}

func readInput() (float64, string, string) {
	fromCurrency := readCurrency("Enter source currency (USD/EUR/RUB): ")
	amount := readAmount()
	toCurrency := readTargetCurrency(fromCurrency)

	return amount, fromCurrency, toCurrency
}

func readCurrency(prompt string) string {
	for {
		fmt.Print(prompt)
		var currency string
		fmt.Scan(&currency)
		if isValidCurrency(currency) {
			return currency
		}
		fmt.Println("Invalid currency. Try again.")
	}
}

func readTargetCurrency(fromCurrency string) string {
	for {
		fmt.Print("Enter target currency (USD/EUR/RUB): ")
		var currency string
		fmt.Scan(&currency)
		if !isValidCurrency(currency) {
			fmt.Println("Invalid currency. Try again.")
			continue
		}
		if currency == fromCurrency {
			fmt.Println("Currencies must be different. Try again.")
			continue
		}
		return currency
	}
}

func readAmount() float64 {
	for {
		fmt.Print("Enter amount: ")
		var amount float64
		_, err := fmt.Scan(&amount)
		if err != nil || amount <= 0 {
			fmt.Println("Invalid amount. Try again.")
			continue
		}
		return amount
	}
}

type currenciesMap map[string]map[string]float64

func calcResult(amount float64, fromCurrency, toCurrency string) float64 {
	currencies := currenciesMap{
		"USD": {
			"EUR": USDToEUR,
			"RUB": USDToRUB,
		},
		"EUR": {
			"USD": EURToUSD,
			"RUB": EURToRUB,
		},
		"RUB": {
			"USD": USDToRUB,
			"EUR": EURToUSD,
		},
	}
	return amount * currencies[fromCurrency][toCurrency]
}
