package main

import (
	"fmt"
	"go-lession/account"
	"go-lession/encrypter"
	"go-lession/files"
	"go-lession/output"
	"strings"

	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": searchAccountByUrl,
	"3": searchAccountByLogin,
	"4": deleteAccount,
}

func main() {
	err := godotenv.Load()
	vault := account.NewVault(files.NewJsonDb("test.vault"), *encrypter.NewEncrypter())
	if err != nil {
		output.PrintError(err)
	}
Menu:
	for {
		choose := promptData(
			"1. Создать аккаунт",
			"2. Найти аккаунт по URL",
			"3. Найти аккаунт по логину",
			"4. Удалить аккаунт",
			"5. Выйти",
			"Выберите пункт меню",
		)

		menuFunc := menu[choose]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}
}

func createAccount(vault *account.VaultWithDb) {
	acc := account.NewAccount()
	vault.AddAccount(*acc)
}

func searchAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите URL")
	data, err := vault.FindAccounts(url, func(a account.Account, s string) bool {
		return strings.Contains(a.Url, s)
	})
	outputResult(&data, err)
}

func searchAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	data, err := vault.FindAccounts(login, func(a account.Account, s string) bool {
		return strings.Contains(a.Login, s)
	})
	outputResult(&data, err)
}

func outputResult(accounts *[]account.Account, err error) {
	if err != nil {
		output.PrintError("Ничего не найдено")
		return
	}
	fmt.Println("\nВот что я нашёл:")
	for _, value := range *accounts {
		value.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для удаления")
	isDelete := vault.DeleteAccountByUrl(url)
	if !isDelete {
		output.PrintError("Ничего не нашлось")
		return
	}
	fmt.Println("Запись успешно удалена")
}

func promptData(prompt ...any) string {
	var res string
	for i, val := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", val)
			continue
		}
		fmt.Println(val)
	}
	fmt.Scan(&res)
	return res
}
