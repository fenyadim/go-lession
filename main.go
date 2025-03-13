package main

import (
	"fmt"
	"go-lession/account"
	"go-lession/files"
	"go-lession/output"
)

func main() {
	vault := account.NewVault(files.NewJsonDb("test.json"))
Menu:
	for {
		choose := inputMenu()
		switch choose {
		case 1:
			createAccount(vault)
		case 2:
			searchAccount(vault)
		case 3:
			deleteAccount(vault)
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

func createAccount(vault *account.VaultWithDb) {
	acc := account.NewAccount()
	vault.AddAccount(*acc)
}

func searchAccount(vault *account.VaultWithDb) {
	var search string
	fmt.Print("Введите URL: ")
	fmt.Scan(&search)
	data, err := vault.FindAccountsByUrl(search)
	if err != nil {
		output.PrintError("Ничего не найдено")
		return
	}
	fmt.Println("\nВот что я нашёл:")
	for _, value := range data {
		value.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	var url string
	fmt.Print("Введите URL для удаления: ")
	fmt.Scan(&url)
	isDelete := vault.DeleteAccountByUrl(url)
	if !isDelete {
		output.PrintError("Ничего не нашлось")
		return
	}
	fmt.Println("Запись успешно удалена")
}
