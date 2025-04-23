package main

import (
	"errors"
	"fmt"
)

const (
	USDtoEUR = 0.92
	USDtoRUB = 84.87
	EURtoRUB = (1 / USDtoEUR) * USDtoRUB
	EURtoUSD = 1 / USDtoEUR
	RUBtoEUR = 1 / EURtoRUB
	RUBtoUSD = 1 / USDtoRUB

	RUB = "RUB"
	USD = "USD"
	EUR = "EUR"

	minNumber = 0.001
	maxNumber = 9999999999999999.0
)

var rateStorage *currencyMap

type currencyMap = map[string]singleCurrMap
type singleCurrMap = map[string]float64

func init() {
	m := make(currencyMap)
	rateStorage = &m
	(*rateStorage)[RUB] = make(singleCurrMap)
	(*rateStorage)[USD] = make(singleCurrMap)
	(*rateStorage)[EUR] = make(singleCurrMap)

	(*rateStorage)[RUB][USD] = RUBtoUSD
	(*rateStorage)[RUB][EUR] = RUBtoEUR
	(*rateStorage)[USD][EUR] = USDtoEUR
	(*rateStorage)[USD][RUB] = USDtoRUB
	(*rateStorage)[EUR][RUB] = EURtoRUB
	(*rateStorage)[EUR][USD] = EURtoUSD

}

func main() {
	fmt.Println("IIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIII")
	fmt.Println("       __Калькулятор валют__")

	for {
		fmt.Printf("Конвертация 1 USD в EUR = %0.3f\n", USDtoEUR)
		fmt.Printf("Конвертация 1 USD в RUB = %0.3f\n", USDtoRUB)
		fmt.Printf("Конвертация 1 EUR в RUB = %0.3f\n", EURtoRUB)
		fmt.Printf("Конвертация 1 EUR в USD = %0.3f\n", EURtoUSD)
		fmt.Printf("Конвертация 1 RUB в EUR = %0.3f\n", RUBtoEUR)
		fmt.Printf("Конвертация 1 RUB в USD = %0.3f\n", RUBtoUSD)

		number, firstCurr, secondCurr, _ := inputNumbers()
		ansver, err := calculationCurr(number, firstCurr, secondCurr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("        Ваш перевод:%0.3f\n", ansver)

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
	fmt.Println("Введите желаемое число для перевода")
	fmt.Print("Например 10, 200, 3000:")
	fmt.Scan(&number)
	fmt.Println("Введите первую желаемую валюту")
	fmt.Print("Например RUB, USD, EUR:")
	fmt.Scan(&firstCurr)
	if firstCurr != "RUB" && firstCurr != "USD" && firstCurr != "EUR" {
		return 0, "0", "0", errors.New("Введите валюту из списка")
	}
	fmt.Println("Введите вторую желаемую валюту")
	fmt.Print("Например RUB, USD, EUR:")
	fmt.Scan(&secondCurr)
	if secondCurr != "RUB" && secondCurr != "USD" && secondCurr != "EUR" {
		return 0, "0", "0", errors.New("Введите валюту из списка")
	}
	return number, firstCurr, secondCurr, nil
}

func calculationCurr(number float64, firstCurr string, secondCurr string) (float64, error) {
	if number <= minNumber || number >= maxNumber || firstCurr == secondCurr {
		return 0, errors.New("Введите правильные значения!\n")
	}

	rate, status := (*rateStorage)[firstCurr][secondCurr]
	if !status {
		return 0, errors.New("Такого курса нет в списке!")
	}

	return rate * number, nil
}

func yesOrNot() bool {
	var userChoise string
	fmt.Println("Хотите ли вы сделать еще расчет?(Y/N)")
	fmt.Scan(&userChoise)
	if userChoise == "Y" || userChoise == "y" {
		return true
	}
	return false
}
