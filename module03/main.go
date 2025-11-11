package main

import (
	"Project_1/module03/main/models"
	"encoding/json"
	"fmt"
)

func main() {
	user1 := models.User{Name: "Иван", Surname: "Иванов", Email: "ivan@example.com", Age: 30}
	user2 := models.User{}
	user3 := models.User{Name: "Alice"}
	user4 := models.User{Name: "Алиса"}
	user5 := models.User{Name: "Алиса", Surname: "Smith"}
	user6 := models.User{Name: "Алиса", Surname: "Солнцева"}
	user7 := models.User{Name: "Алиса", Surname: "Солнцева", Email: "alice"}

	users := []models.User{user1, user2, user3, user4, user5, user6, user7}

	for i, u := range users {
		if err := u.Validate(); err != nil {
			fmt.Printf("Ошибка валидации для пользователя №%v: %s\n", i+1, err)
		} else {
			fmt.Printf("Пользователь №%v указан корректно.\n", i+1)
		}
	}

	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("Ошибка сериализации:", err)
		return
	}

	fmt.Println("JSON:", string(data))
	var parsedUser models.User
	err = json.Unmarshal(data, &parsedUser)
	if err != nil {
		fmt.Println("Ошибка десериализации:", err)
		return
	}
	fmt.Printf("Parsed User: %+v\n", parsedUser)
}
