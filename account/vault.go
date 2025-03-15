package account

import (
	"encoding/json"
	"errors"
	"go-lession/encrypter"
	"go-lession/output"
	"slices"
	"strings"
	"time"

	"github.com/fatih/color"
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
	db  Db
	enc encrypter.Encrypter
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
	encData := vault.enc.Encrypt(data)
	if err != nil {
		output.PrintError(err)
	}
	vault.db.Write(encData)
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}

	data := enc.Decrypt(file)

	var vault Vault
	err = json.Unmarshal(data, &vault)
	color.Cyan("Найдено %d аккаунтов", len(vault.Accounts))
	if err != nil {
		output.PrintError(err)
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}
