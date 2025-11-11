package models

import (
	"errors"
	"regexp"
)

type User struct {
	Name    string
	Surname string
	Email   string
	Age     uint
}

func (user User) Validate() error {
	if user.Name == "" {
		return errors.New("Имя не может быть пустым.")
	}

	regex := regexp.MustCompile("^[а-яёА-ЯЁ]*$")
	if !regex.MatchString(user.Name) {
		return errors.New("Имя содержит некорректные символы. Введите имя, используя только русские буквы.")
	}

	if user.Surname == "" {
		return errors.New("Фамилия не может быть пустой.")
	}

	if !regex.MatchString(user.Surname) {
		return errors.New("Фамилия содержит некорректные символы. Введите фамилию, используя только русские буквы.")
	}

	if user.Email == "" {
		return errors.New("Email должен быть задан.")
	}

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(user.Email) {
		return errors.New("Некорректный email.")
	}

	return nil
}
