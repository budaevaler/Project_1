package main

import (
	"errors"
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(fibonachi(6))

	fmt.Println(sumInt(1, 4, 5, 6))

	a := 7
	b := 4

	multiply(&a, &b)

	swap(a, b)

	fmt.Println(divide(a, b))

	fmt.Println(validate("12345Avdd"))
	fmt.Println(validate("AAAAA"))
	fmt.Println(validate("смир12"))
	fmt.Println(validate("смир2"))

}

func fibonachi(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonachi(n-1) + fibonachi(n-2)
}

func sumInt(numbers ...int) (int, int) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return len(numbers), sum
}

func multiply(a *int, b *int) {
	fmt.Println((*a) * (*b))
}

func swap(a int, b int) {
	fmt.Printf("before swap a=%v\r\n", a)
	fmt.Printf("before swap b=%v\r\n", b)
	pa := &a
	pb := &b
	fmt.Printf("before swap pointer a=%v\r\n", pa)
	fmt.Printf("before swap pointer b=%v\r\n", pb)
	*pa, *pb = *pb, *pa
	fmt.Printf("after swap pointer a=%v\r\n", pa)
	fmt.Printf("after swap pointer b=%v\r\n", pb)
	fmt.Printf("after swap a=%v\r\n", a)
	fmt.Printf("after swap b=%v\r\n", b)
}

func divide(a int, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("my generated error")
	}
	res = a / b
	return
}

func validate(password string) string {
	re := regexp.MustCompile("^[a-zA-Z0-9]*$")
	var res string
	if len([]rune(password)) >= 5 && re.MatchString(password) {
		res = "Ok"
	} else {
		res = "Wrong password"
	}
	return res
}
