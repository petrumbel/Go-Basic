package main

import (
	"fmt"
)

const USDtoEUR = 0.92
const USDtoRUB = 84.87
const EURtoRUB = (1 / USDtoEUR) * USDtoRUB
const EURtoUSD = 1 / USDtoEUR
const RUBtoEUR = 1 / EURtoRUB
const RUBtoUSD = 1 / USDtoRUB

func main() {
	fmt.Println("IIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIII")
	fmt.Println("|      __Калькулятор валют__          |")
	fmt.Printf("|Конвертация 1 USD в EUR = %0.3f      |\n", USDtoEUR)
	fmt.Printf("|Конвертация 1 USD в RUB = %0.3f     |\n", USDtoRUB)
	fmt.Printf("|Конвертация 1 EUR в RUB = %0.3f     |\n", EURtoRUB)
	fmt.Printf("|Конвертация 1 EUR в USD = %0.3f      |\n", EURtoUSD)
	fmt.Printf("|Конвертация 1 RUB в EUR = %0.3f      |\n", RUBtoEUR)
	fmt.Printf("|Конвертация 1 RUB в USD = %0.3f      |\n", RUBtoUSD)

	number, firstCurr, secondCurr := inputNumbers()
	ansver := calculation(number, firstCurr, secondCurr)
	fmt.Printf("|     Ваш перевод:%0.3f\n", ansver)
	fmt.Println("IIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIII")

}

func inputNumbers() (float64, string, string) {
	var number float64
	var firstCurr string
	var secondCurr string
	fmt.Println("|Введите желаемое число для перевода  |")
	fmt.Print("|Например 10, 200, 3000:")
	fmt.Scan(&number)
	fmt.Println("|                                     |")
	fmt.Println("|Введите первую желаемую валюту       |")
	fmt.Print("|Например RUB, USD, EUR:")
	fmt.Scan(&firstCurr)
	fmt.Println("|                                     |")
	fmt.Println("|Введите вторую желаемую валюту       |")
	fmt.Print("|Например RUB, USD, EUR:")
	fmt.Scan(&secondCurr)
	fmt.Println("|                                     |")
	return number, firstCurr, secondCurr
}

func calculation(number float64, firstCurr string, secondCurr string) float64 {
	var ansver float64
	if firstCurr == "RUB" || firstCurr == "rub" && secondCurr == "USD" || secondCurr == "usd" {
		ansver = number * RUBtoUSD
	} else if firstCurr == "RUB" || firstCurr == "rub" && secondCurr == "EUR" || secondCurr == "eur" {
		ansver = number * RUBtoEUR
	} else if firstCurr == "USD" || firstCurr == "usd" && secondCurr == "RUB" || secondCurr == "rub" {
		ansver = number * USDtoRUB
	} else if firstCurr == "USD" || firstCurr == "usd" && secondCurr == "EUR" || secondCurr == "eur" {
		ansver = number * USDtoEUR
	} else if firstCurr == "EUR" || firstCurr == "eur" && secondCurr == "USD" || secondCurr == "usd" {
		ansver = number * EURtoUSD
	} else if firstCurr == "EUR" || firstCurr == "eur" && secondCurr == "RUB" || secondCurr == "rub" {
		ansver = number * EURtoRUB
	}
	return ansver

}

// Финализируем приложение калькулятор. Для этого:
// - Сделать меню с шагами
// - Ввод исходной валюты (подсказываем варианты) -
// если ошибка, заново вводим
// - Ввод числа - если ошибка, заново вводим
// - Ввод целевой валюты (подсказываем варианты) - если
// ошибка, заново вводим
// Выделить функцию ввода / проверки валюты и числа
// После получения всех данных с помощью if / switch вычислить
// итог и вывести результат.
