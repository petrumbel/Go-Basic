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
	if err := ensureDataFile(); err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}

	binList, err := loadBins()
	if err != nil {
		fmt.Println("Ошибка при загрузке bin:", err)
		return
	}

	newBin := promptBinInput()
	binList = append(binList, newBin)

	if err := saveBinList(binList); err != nil {
		fmt.Println("Ошибка при сохранении bin:", err)
		return
	}

	fmt.Println("Bin успешно создан и сохранён!")
}

func ensureDataFile() error {
	return file.EnsureFileExists("data.json")
}

func loadBins() (bins.BinList, error) {
	return storage.LoadBins()
}

func promptBinInput() bins.Bin {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите ID: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	fmt.Print("Введите Логин: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Приватный? (y/n): ")
	privateInput, _ := reader.ReadString('\n')
	privateInput = strings.TrimSpace(strings.ToLower(privateInput))
	private := privateInput == "y" || privateInput == "yes"

	return bins.NewBin(id, name, private)
}

func saveBinList(binList bins.BinList) error {
	return storage.SaveBins(binList)
}
