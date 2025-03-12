package account

import (
	"encoding/json"
	"math/rand"
	"time"
)

var letterRuns = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890-")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) generatePassword(num int) {
	arr := make([]rune, num)
	for i := range arr {
		arr[i] = letterRuns[rand.Intn(len(letterRuns))]
	}
	acc.Password = string(arr)
}

func (acc *Account) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewAccount() *Account {
	newAcc := &Account{
		Login:     "Test",
		Password:  "",
		Url:       "https://test.ru",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if newAcc.Password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc
}
