package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-lession/files"
	"slices"
	"strings"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) FindAccountsByUrl(url string) ([]Account, error) {
	res := []Account{}
	for _, value := range vault.Accounts {
		if strings.Contains(value.Url, url) {
			res = append(res, value)
		}
	}

	if len(res) == 0 {
		err := errors.New("NOT_FOUND")
		return nil, err
	}

	return res, nil
}

func (vault *Vault) DeleteAccountByUrl(url string) bool {
	isDeleted := false
	for index, account := range vault.Accounts {
		if strings.Contains(account.Url, url) {
			vault.Accounts = slices.Delete(vault.Accounts, index, index+1)
			isDeleted = true
		}
	}
	vault.save()
	return isDeleted
}

func (vault *Vault) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println(err.Error())
	}
	files.WriteFile(data, "test.json")
}

func NewVault() *Vault {
	file, err := files.ReadFile("test.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &vault
}
