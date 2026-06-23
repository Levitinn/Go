package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operation := readOperation()
	numbers := readNumbers()
	result := executeOperation(operation, numbers)
	fmt.Printf("%.2f\n", result)
}

func isValidOperation(operation string) bool {
	return operation == "AVG" || operation == "SUM" || operation == "MED"
}
func readOperation() string {

	fmt.Println("Введите операцию: AVG/SUM/MED")
	var operation string
	fmt.Scan(&operation)
	if !isValidOperation(operation) {
		fmt.Println("Неверная операция. Попробуйте снова.")
	}
	return operation
}

func readNumbers() []float64 {
	numbers := []float64{}
	fmt.Println("Введите числа через запятую:")
	input := ""
	fmt.Scanln(&input)
	parts := strings.Split(input, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		num, err := strconv.ParseFloat(part, 64)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}
func executeOperation(operation string, numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	switch operation {
	case "SUM":
		sum := 0.0
		for _, num := range numbers {
			sum += num
		}
		return sum

	case "AVG":
		sum := 0.0
		for _, num := range numbers {
			sum += num
		}
		return sum / float64(len(numbers))

	case "MED":
		sorted := make([]float64, len(numbers))
		for i, num := range numbers {
			sorted[i] = num
		}
		sort.Float64s(sorted)

		mid := len(sorted) / 2
		if len(sorted)%2 == 1 {
			return sorted[mid]
		}
		return (sorted[mid-1] + sorted[mid]) / 2
	}

	return 0
}
