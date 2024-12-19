package repository

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type JsonRepository struct {
	filename string
}

func NewJsonRepository(name string) Repository {
	return &JsonRepository{
		filename: name,
	}
}

func (jsonRepo *JsonRepository) Read() ([]byte, error) {
	bytes, err := os.ReadFile(jsonRepo.filename)

	if err != nil {
		color.Red(err.Error())
		return nil, err
	}

	return bytes, nil
}

func (jsonRepo *JsonRepository) Write(content []byte) {
	file, err := os.Create(jsonRepo.filename)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	if _, err = file.Write(content); err != nil {
		fmt.Println(err)
		return
	}

	color.Green("Запись успешна")
}
