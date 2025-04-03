package main

import "fmt"

const USDtoEUR = 0.92
const USDtoRUB = 84.87
const EURtoRUB = (1 / USDtoEUR) * USDtoRUB
const EURtoUSD = 1 / USDtoEUR
const RUBtoEUR = 1 / EURtoRUB
const RUBtoUSD = 1 / USDtoRUB

func main() {
	fmt.Printf("Конвертация 1 USD в EUR = %0.3f\n", USDtoEUR)
	fmt.Printf("Конвертация 1 USD в RUB = %0.3f\n", USDtoRUB)
	fmt.Printf("Конвертация 1 EUR в RUB = %0.3f\n", EURtoRUB)
	fmt.Printf("Конвертация 1 EUR в USD = %0.3f\n", EURtoUSD)
	fmt.Printf("Конвертация 1 RUB в EUR = %0.3f\n", RUBtoEUR)
	fmt.Printf("Конвертация 1 RUB в USD = %0.3f\n", RUBtoUSD)
	number, firstCurr, secondCurr := inputNumbers()
	ansver := calculation(number, firstCurr, secondCurr)
	fmt.Printf("%0.3f", ansver)
}

func inputNumbers() (float64, string, string) {
	var number float64
	var firstCurr string
	var secondCurr string
	fmt.Scan(&number)
	fmt.Scan(&firstCurr)
	fmt.Scan(&secondCurr)
	return number, firstCurr, secondCurr
}

func calculation(number float64, firstCurr string, secondCurr string) float64 {
	var ansver float64
	if firstCurr == "RUB" && secondCurr == "USD" {
		ansver = number * RUBtoUSD
	} else {
		fmt.Print("ничего не вышло")
	}
	return ansver

}

//Сделать считывание ввода пользователя в виде отдельной
//функции
//- Сделать пустую функцию расчета, которая будет принимать 1
//число и 2 валюты - исходная и целевая
