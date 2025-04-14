package main

import (
	"errors"
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
	for {
		fmt.Printf("|Конвертация 1 USD в EUR = %0.3f      |\n", USDtoEUR)
		fmt.Printf("|Конвертация 1 USD в RUB = %0.3f     |\n", USDtoRUB)
		fmt.Printf("|Конвертация 1 EUR в RUB = %0.3f     |\n", EURtoRUB)
		fmt.Printf("|Конвертация 1 EUR в USD = %0.3f      |\n", EURtoUSD)
		fmt.Printf("|Конвертация 1 RUB в EUR = %0.3f      |\n", RUBtoEUR)
		fmt.Printf("|Конвертация 1 RUB в USD = %0.3f      |\n", RUBtoUSD)
		number, firstCurr, secondCurr, _ := inputNumbers()
		ansver, err := calculationCurr(number, firstCurr, secondCurr)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("|       Ваш перевод:%0.3f\n", ansver)
		}
		checkRepeate := yesOrNot()
		if !checkRepeate {
			break
		}
		fmt.Println("IIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIII")
	}
}

func inputNumbers() (float64, string, string, error) {
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
	if firstCurr != "RUB" && firstCurr != "USD" && firstCurr != "EUR" {
		return 0, "0", "0", errors.New("Введите валюту из списка")
	}
	fmt.Println("|                                     |")
	fmt.Println("|Введите вторую желаемую валюту       |")
	fmt.Print("|Например RUB, USD, EUR:")
	fmt.Scan(&secondCurr)
	if secondCurr != "RUB" && secondCurr != "USD" && secondCurr != "EUR" {
		return 0, "0", "0", errors.New("Введите валюту из списка")
	}
	fmt.Println("|                                     |")
	return number, firstCurr, secondCurr, nil
}

func calculationCurr(number float64, firstCurr string, secondCurr string) (float64, error) {
	if number >= 0.001 && number <= 999999999999999.0 && firstCurr != secondCurr {
		switch {
		case firstCurr == "RUB" && secondCurr == "EUR":
			return number * RUBtoEUR, nil
		case firstCurr == "RUB" && secondCurr == "USD":
			return number * RUBtoUSD, nil
		case firstCurr == "EUR" && secondCurr == "USD":
			return number * EURtoUSD, nil
		case firstCurr == "EUR" && secondCurr == "RUB":
			return number * EURtoRUB, nil
		case firstCurr == "USD" && secondCurr == "RUB":
			return number * USDtoRUB, nil
		case firstCurr == "USD" && secondCurr == "EUR":
			return number * USDtoEUR, nil
		}
	} else {
		return 0, errors.New("Введите правильные значения!")
	}
	return 0, nil
}

func yesOrNot() bool {
	var userChoise string
	fmt.Println("|Хотите ли вы сделать еще расчет?(Y/N)|")
	fmt.Scan(&userChoise)
	if userChoise == "Y" || userChoise == "y" {
		return true
	}
	return false
}
