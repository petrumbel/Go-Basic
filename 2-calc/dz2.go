package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("__Калькулятор чисел__")
	operation, stringNumbers := input()
	arrayNumbers, err := parseNumbers(stringNumbers)
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
	}
	choosingAnOperation(operation, arrayNumbers)

}

func input() (string, string) {
	var operation string
	var stringNumbers string

	fmt.Print("Введите операцию(AVG, SUM, MED): ")
	fmt.Scan(&operation)

	fmt.Print("Введите числа через запятую: ")
	fmt.Scan(&stringNumbers)

	return operation, stringNumbers
}

func choosingAnOperation(operation string, arrayNumbers []float64) {
	switch operation {
	case "AVG":
		AVGOperation(arrayNumbers)
	case "SUM":
		SUMOperation(arrayNumbers)
	case "MED":
		MEDOperation(arrayNumbers)
	}

}

func parseNumbers(stringNumbers string) ([]float64, error) {
	parts := strings.Split(stringNumbers, ",")
	numbers := make([]float64, 0, len(parts))

	for _, part := range parts {
		trimmedString := strings.TrimSpace(part)
		if trimmedString == "" {
			continue
		}

		number, err := strconv.ParseFloat(trimmedString, 64)
		if err != nil {
			return nil, fmt.Errorf("некорректное число: %s", trimmedString)

		}
		numbers = append(numbers, number)

	}
	if len(numbers) == 0 {
		return nil, fmt.Errorf("нет чисел для обработки")
	}

	return numbers, nil
}

func AVGOperation(arrayNumbers []float64) {
	fmt.Println("Входные данные:", arrayNumbers)
	sum := 0.0
	for _, num := range arrayNumbers {
		sum += num
	}
	avg := sum / float64(len(arrayNumbers))

	fmt.Printf("AVG = %0.2f", avg)
}

func SUMOperation(arrayNumbers []float64) {
	sum := 0.0
	for _, num := range arrayNumbers {
		sum += num
	}
	fmt.Printf("SUM = %0.2f", sum)
}

func MEDOperation(arrayNumbers []float64) {
	data := make([]float64, len(arrayNumbers))
	copy(data, arrayNumbers)

	sort.Float64s(data)

	n := len(data)
	var median float64

	if n == 0 {
		fmt.Print("Нет данных для вычисления медианы")
		return
	}

	if n%2 == 1 {
		median = data[n/2]
	} else {
		median = (data[n/2-1] + data[n/2]) / 2
	}

	fmt.Printf("MED = %0.2f", median)
}
