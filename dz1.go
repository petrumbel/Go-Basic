package main

import "fmt"

func main() {
	const USDtoEUR = 0.92
	const USDtoRUB = 84.87
	const EURtoRUB = (1 / USDtoEUR) * USDtoRUB
	fmt.Printf("Конвертация 1 USD в EUR = %v\n", USDtoEUR)
	fmt.Printf("Конвертация 1 USD в RUB = %v\n", USDtoRUB)
	fmt.Printf("Конвертация 1 USD в RUB = %v\n", EURtoRUB)

}
