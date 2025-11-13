package main

import (
	"Project_1/module08tasks/main/models"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//task1()
	//task2()
	task3()
}

func task1() {
	weather := models.Weather{Date: time.Date(2025, 11, 13, 15, 40, 0, 0, time.UTC), Temperature: 2, Description: "windy"}

	weatherJson, err := json.Marshal(weather)
	if err != nil {
		fmt.Printf("При сериализации возникла ошибка: ", err)
	}

	var weatherCopy models.Weather
	err = json.Unmarshal(weatherJson, &weatherCopy)
	if err != nil {
		fmt.Printf("При десериализации возникла ошибка: ", err)
	}

	fmt.Println(weatherCopy)
}

func task2() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите 2 даты через запятую..")

	//Thu Nov 13 15:57:48 MSK 2025,14-11-2025 16:07:08
	//Thu Nov 13 15:57:48 MSK 2025,10-11-2025 16:07:08
	var inputString string
	if scanner.Scan() {
		inputString = scanner.Text()
		inputString = strings.TrimSpace(inputString)
	} else {
		fmt.Println("Ошибка при чтении строки")
		return
	}

	dates := strings.Split(inputString, ",")
	if len(dates) != 2 {
		fmt.Println("Введено неверное количество дат.")
		return
	}

	diff, err := getDiffBetweenDates(dates[0], dates[1])
	if err != nil {
		fmt.Println("При выполнении задания 2 возникла ошибка:", err)
		return
	}

	fmt.Println(diff)
}

func getDiffBetweenDates(dateString1 string, dateString2 string) (duration time.Duration, parseError error) {
	time1, err := time.Parse("Mon Jan 2 15:04:05 MST 2006", dateString1)
	if err != nil {
		parseError = err
		return
	}

	time2, err := time.Parse("02-01-2006 15:04:05", dateString2)
	if err != nil {
		parseError = err
		return
	}

	if time1.After(time2) {
		duration = time1.Sub(time2)
	} else {
		duration = time2.Sub(time1)
	}

	return
}

func task3() {
	fmt.Printf("Введите число. Для завершения ввода введите букву.")
	scanner := bufio.NewScanner(os.Stdin)

	var numbers []byte
	var sum int
	sum = 0

	for scanner.Scan() {
		line := scanner.Text()
		if n, err := strconv.ParseInt(line, 10, 8); err != nil {
			if n > 100 || n < 0 {
				continue
			}
			numbers = append(numbers, byte(n))
			sum += int(n)
		} else {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}

	fmt.Println("Введены следующие числа:")
	for _, v := range numbers {
		fmt.Println(v)
	}

	fileName := "sum.txt"

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Ошибка открытия файла %s: %v\n", fileName, err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(sum) + "\n")
	if err != nil {
		fmt.Printf("Ошибка записи в файл: %v\n", err)
		return
	}
}
