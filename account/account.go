package account

import (
	"errors"
	"math/rand"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letters []rune = []rune("qwerrtyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if err := validateAccountData(login, urlString); err != nil {
		return nil, errors.New(err.Error())
	}

	account := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if password == "" {
		account.generatePassword(10)
	}

	return account, nil
}

func validateAccountData(login, urlString string) error {
	if _, err := url.ParseRequestURI(urlString); err != nil {
		return errors.New("INVALID_URL")
	}

	if login == "" {
		return errors.New("EMPTY_LOGIN")
	}

	return nil
}

func (account *Account) generatePassword(length int) string {
	res := make([]rune, length)

	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}

	return string(res)
}

func (account *Account) LogAccountInfo() {
	color.Cyan(account.Login)
	color.Cyan(account.Password)
	color.Cyan(account.Url)
}
