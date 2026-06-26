package main

import (
	"fmt"
	"os"
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
	for {
		fmt.Println("Введите операцию: AVG/SUM/MED")
		var operation string
		fmt.Scanln(&operation)
		if !isValidOperation(operation) {
			fmt.Println("Неверная операция. Попробуйте снова.")
			continue
		}
		return operation
	}
}

func readNumbers() []float64 {
	numbers := []float64{}
	fmt.Println("Введите числа через запятую:")
	var input string
	fmt.Fscanf(os.Stdin, "%[^\n]", &input)

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
	operations := map[string]func([]float64) float64{
		"SUM": func(numbers []float64) float64 {
			sum := 0.0
			for _, num := range numbers {
				sum += num
			}
			return sum
		},
		"AVG": func(numbers []float64) float64 {
			sum := 0.0
			for _, num := range numbers {
				sum += num
			}
			return sum / float64(len(numbers))
		},
		"MED": func(numbers []float64) float64 {
			sorted := make([]float64, len(numbers))
			copy(sorted, numbers)
			sort.Float64s(sorted)

			mid := len(sorted) / 2
			if len(sorted)%2 == 1 {
				return sorted[mid]
			}
			return (sorted[mid-1] + sorted[mid]) / 2
		},
	}
	execute, ok := operations[operation]
	if !ok {
		return 0
	}
	return execute(numbers)
}
