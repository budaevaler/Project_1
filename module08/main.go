package main

import (
	"Project_1/module08/main/models"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("users.csv")
	if err != nil {
		log.Fatal("Ошибка открытия файла:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	_, err = reader.Read()
	if err != nil && err != io.EOF {
		log.Fatal("Ошибка чтения заголовка:", err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Ошибка чтения строки:", err)
		}
		person, err := ParseUser(record)
		if err != nil {
			log.Printf("Ошибка парсинга строки: %v, строка: %v", err, record)
			continue
		}
		fmt.Printf("Прочитано: %+v\n", person)
	}
}

func ParseUser(line []string) (user *models.User, parseError error) {
	if len(line) != 4 {
		return nil, fmt.Errorf("ожидается 4 поля, получено %d", len(line))
	}

	name := line[0]
	surname := line[1]
	email := line[2]
	age, err := strconv.ParseUint(line[3], 10, 64)
	if err != nil {
		parseError = fmt.Errorf("Не удалось прочитать Age: %s из-за ошибки %s", line[3], err)
	}

	user = &models.User{Name: name, Surname: surname, Email: email, Age: uint(age)}
	return
}
