package main

import "fmt"

func main() {
	const USDtoEUR = 0.92
	const USDtoRUB = 84.87
	const EURtoRUB = (1 / USDtoEUR) * USDtoRUB
	fmt.Printf("Конвертация 1 USD в EUR = %v\n", USDtoEUR)
	fmt.Printf("Конвертация 1 USD в RUB = %v\n", USDtoRUB)
	fmt.Printf("Конвертация 1 EUR в RUB = %v\n", EURtoRUB)
	firstCurr, secondCurr, number := inputNumbers()
	calculation(firstCurr, secondCurr, number)
}

//почему в первые ковычки нельзя написать firstCurr и тд.
func inputNumbers() (float64, float64, float64) {
	var firstCurr float64
	var secondCurr float64
	var number float64
	fmt.Scan(&firstCurr)
	fmt.Scan(&secondCurr)
	fmt.Scan(&number)
	return firstCurr, secondCurr, number
}

func calculation(firstCurr, secondCurr, number float64) {

}

//Сделать считывание ввода пользователя в виде отдельной
//функции
//- Сделать пустую функцию расчета, которая будет принимать 1
//число и 2 валюты - исходная и целевая
