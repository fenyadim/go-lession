package main

import (
	"fmt"
	"go-lession/account"
	"go-lession/files"
)

func main() {
Menu:
	for {
		choose := inputMenu()
		switch choose {
		case 1:
			createAccount()
		case 2:
			searchAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func inputMenu() int {
	var input int
	fmt.Println("Выберите пункт меню")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выйти")
	fmt.Scan(&input)
	return input
}

func createAccount() {
	acc := account.NewAccount()
	file, err := acc.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}
	files.WriteFile(file, "test.json")
	files.ReadFile()
}

func searchAccount() {

}

func deleteAccount() {

}
