package main

import (
	account "alari/passwordGeneration/account"
	repository "alari/passwordGeneration/repository"
	"alari/passwordGeneration/service"
	utils "alari/passwordGeneration/utils"
	"fmt"

	color "github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.AccountStore){
	"1": createAccount,
	"2": findAccount,
	"3": deleteAccount,
}

func main() {
	err := godotenv.Load()
	if err != nil {
		color.Red("Не удалось загрузить env переменные из .env файла")
	}

	fmt.Println("__Менеджер паролей__")
	accountStore := account.NewAccountStore(repository.NewJsonRepository("data.vault"), *service.NewEncrypter())

Menu:
	for {
		variant := getVariant()

		menuFunc := menu[variant]

		if menuFunc == nil {
			break Menu
		}

		menuFunc(accountStore)
	}
}

func createAccount(accountStore *account.AccountStore) {
	login := utils.PromptData("Введите логин: ")
	password := utils.PromptData("Введите пароль: ")
	urlString := utils.PromptData("Введите url: ")

	newAccount, err := account.NewAccount(login, password, urlString)

	if err != nil {
		panic(err.Error())
	}

	accountStore.AddAccount(*newAccount)
}

func findAccount(accountStore *account.AccountStore) {
	url := utils.PromptData("Введите url")

	findedAccountsByUrl := accountStore.FindAccountByUrl(url)

	if len(findedAccountsByUrl) == 0 {
		color.Red("Не одного аккаунта не было найдено")
		return
	}

	for _, account := range findedAccountsByUrl {
		account.LogAccountInfo()
	}
}

func deleteAccount(accountStore *account.AccountStore) {
	url := utils.PromptData("Введите url")

	wasSomeAccountsDeleted := accountStore.DeleteAccount(url)

	if wasSomeAccountsDeleted {
		color.Green("Аккаунт успешно удален")
	} else {
		color.Red("Не найдено")
	}
}

func getVariant() string {
	var variant string

	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)

	return variant
}
