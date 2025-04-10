package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("__Калькулятор чисел__")
	operation := inputOperation()
	numbers := inputNumbers()
	parseNumbers(numbers)
	arrayNumbers, err := parseNumbers(numbers)
	AVGAnsver := AVGOperation(arrayNumbers)
	SUMAnsver := SUMOperation(arrayNumbers)
	MEDAnsver := MEDOperation(arrayNumbers)
	choosingAnOperation(operation, arrayNumbers)
	fmt.Printf("%v", AVGAnsver)
	fmt.Printf("%v", err)
	fmt.Printf("%v", SUMAnsver)
	fmt.Printf("%v", MEDAnsver)

}

func inputOperation() string {
	var operation string
	fmt.Print("Введите операцию(AVG, SUM, MED): ")
	fmt.Scanln(&operation)
	return operation
}

func inputNumbers() string {
	var stringNumbers string
	fmt.Print("Введите числа через запятую: ")
	fmt.Scanln(&stringNumbers)
	return stringNumbers
}

func choosingAnOperation(operation string, arrayNumbers []float64) string {
	switch operation {
	case "AVG":
		AVGOperation(arrayNumbers)
	case "SUM":
		SUMOperation(arrayNumbers)
	case "MED":
		MEDOperation(arrayNumbers)
	}
	return operation
}

func parseNumbers(numbers string) ([]float64, error) {
	parts := strings.Split(numbers, ",")
	var arrayNumbers []float64

	for _, part := range parts {
		numStr := strings.TrimSpace(part)
		if numStr == "" {
			continue
		}
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return nil, fmt.Errorf("'%s' не является числом", numStr)
		}
		arrayNumbers = append(arrayNumbers, num)
	}

	if len(arrayNumbers) == 0 {
		return nil, fmt.Errorf("не введено ни одного числа")
	}

	return arrayNumbers, nil
}

func AVGOperation(arrayNumbers []float64) float64 {
	sum := 0.0
	for _, num := range arrayNumbers {
		sum += num
	}
	return sum / float64(len(arrayNumbers))

}

func SUMOperation(arrayNumbers []float64) float64 {
	sum := 0.0
	for _, num := range arrayNumbers {
		sum += num

	}
	return sum
}

func MEDOperation(arrayNumbers []float64) float64 {
	sorted := make([]float64, len(arrayNumbers))
	copy(sorted, arrayNumbers)
	n := len(sorted)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	if n%2 == 1 {
		return sorted[n/2]
	}
	return (sorted[n/2-1] + sorted[n/2]) / 2
}
