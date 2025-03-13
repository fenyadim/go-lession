package account

import (
	"encoding/json"
	"errors"
	"go-lession/output"
	"slices"
	"strings"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db Db
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) ([]Account, error) {
	res := []Account{}
	for _, value := range vault.Accounts {
		if checker(value, str) {
			res = append(res, value)
		}
	}

	if len(res) == 0 {
		err := errors.New("NOT_FOUND")
		return nil, err
	}

	return res, nil
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
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

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		output.PrintError(err)
	}
	vault.db.Write(data)
}

func NewVault(db Db) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		output.PrintError(err)
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
	}
}
