package repository

import "fmt"

type AwsRepository struct {
	url string
}

func NewAwsRepository(url string) Repository {
	return &AwsRepository{
		url: url,
	}
}

func (awsRepo AwsRepository) Read() ([]byte, error) {
	fmt.Println("Read from AWS")
	return []byte{}, nil
}

func (awsRepo AwsRepository) Write(content []byte) {
	fmt.Println("Write to AWS")
}
