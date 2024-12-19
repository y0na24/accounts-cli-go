package account

import (
	"alari/passwordGeneration/repository"
	"encoding/json"
	"strings"
	"time"

	color "github.com/fatih/color"
)

type AccountStore struct {
	Accounts   []Account
	UpdatedAt  time.Time
	repository repository.Repository
}

func NewAccountStore(repository repository.Repository) *AccountStore {
	store := &AccountStore{
		Accounts:   []Account{},
		UpdatedAt:  time.Now(),
		repository: repository,
	}

	file, err := store.repository.Read()

	if err != nil {
		return store
	}

	if err = json.Unmarshal(file, store); err != nil {
		color.Red(err.Error())
		return store
	}

	return store
}

func (store *AccountStore) FindAccountByUrl(url string) []Account {
	findedAccounts := make([]Account, 0, 5)

	for _, account := range store.Accounts {
		if isMatched := strings.Contains(account.Url, url); isMatched {
			findedAccounts = append(findedAccounts, account)
		}
	}

	return findedAccounts
}

func (store *AccountStore) AddAccount(account Account) {
	store.Accounts = append(store.Accounts, account)
	store.UpdatedAt = time.Now()

	data, err := store.ToBytes()

	if err != nil {
		color.Red(err.Error())
		return
	}

	store.repository.Write(data)
}

func (store *AccountStore) DeleteAccount(url string) bool {
	accounts := make([]Account, 0, 5)
	wasDeleted := false

	for _, account := range store.Accounts {
		if isMatched := strings.Contains(account.Url, url); !isMatched {
			accounts = append(accounts, account)
			wasDeleted = true
		}
	}

	store.Accounts = accounts
	store.UpdatedAt = time.Now()

	data, err := store.ToBytes()

	if err != nil {
		color.Cyan("Ошибка при переводе store в массив рун")
	}

	store.repository.Write(data)

	return wasDeleted
}

func (store *AccountStore) ToBytes() ([]byte, error) {
	file, err := json.Marshal(store)

	if err != nil {
		return nil, err
	}

	return file, nil
}
