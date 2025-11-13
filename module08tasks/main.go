package main

import (
	"Project_1/module08tasks/main/models"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	task1()

	fmt.Println(task2("Thu Nov 13 15:57:48 MSK 2025", "14-11-2025 16:07:08"))
	fmt.Println(task2("Thu Nov 13 15:57:48 MSK 2025", "10-11-2025 16:07:08"))
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

func task2(dateString1 string, dateString2 string) (duration time.Duration, parseError error) {
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
