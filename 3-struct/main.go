package main

import (
	"bincli/bins"
	"bincli/file"
	"bincli/storage"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if err := file.EnsureFileExists("data.json"); err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}

	binList, err := storage.LoadBins()
	if err != nil {
		fmt.Println("Ошибка при чтении bin:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите ID: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	fmt.Print("Введите логин: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Приватный? (y/n): ")
	privateInput, _ := reader.ReadString('\n')
	privateInput = strings.TrimSpace(strings.ToLower(privateInput))
	private := privateInput == "y" || privateInput == "yes"

	newBin := bins.NewBin(id, name, private)
	binList = append(binList, newBin)

	if err := storage.SaveBins(binList); err != nil {
		fmt.Println("Ошибка при сохранении bin:", err)
		return
	}

	fmt.Println("Bin успешно создан и сохранён!")
}
