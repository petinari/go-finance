package util

import "github.com/brianvoe/gofakeit/v6"

func GetRandomEmail() string {

	return gofakeit.Email()
}

func GetRandomName() string {
	return gofakeit.Name()
}

func GetRandomPassword() string {
	return gofakeit.Password(true, true, true, true, false, 12)
}
