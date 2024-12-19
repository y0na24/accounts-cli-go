package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypter struct {
	Secret string
}

func NewEncrypter() *Encrypter {
	secret := os.Getenv("SECRET")

	if secret == "" {
		panic("Не передан SECRET параметр в env переменные")
	}

	return &Encrypter{
		Secret: secret,
	}
}

func (encrypter *Encrypter) Encrypt(plaingStr []byte) []byte {
	block, err := aes.NewCipher([]byte(encrypter.Secret))

	if err != nil {
		panic(err.Error())
	}

	aesGSM, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGSM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)

	if err != nil {
		panic(err.Error())
	}
	return aesGSM.Seal(nonce, nonce, plaingStr, nil)
}

func (encrypter *Encrypter) Decrypt(encryptedStr []byte) []byte {
	block, err := aes.NewCipher([]byte(encrypter.Secret))

	if err != nil {
		panic(err.Error())
	}

	aesGSM, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGSM.NonceSize()
	nonce, cipherText := encryptedStr[nonceSize:], encryptedStr[:nonceSize]
	plainText, err := aesGSM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}

	return plainText
}
