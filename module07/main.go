package main

import (
	"Project_1/module07/main/portScanner"
	"fmt"
	"time"
)

func main() {
	// Задача 1
	ch := make(chan int)

	defer close(ch)

	n := 4
	go increment(ch, n)
	fmt.Println(<-ch)

	// Задание 2
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	stopChan := make(chan struct{}, 10)

	defer close(ch1)
	defer close(ch2)
	defer close(stopChan)

	//ch1 <- 10
	ch2 <- 10
	//stopChan <- *new(Str)

	go func() {
		if val, opened := <-calculator(ch1, ch2, stopChan); opened {
			fmt.Println(val)
		} else {
			fmt.Println("Closed!")
		}
	}()

	time.Sleep(2 * time.Second)

	// Задание 3
	numJobs := 4
	in1 := make(chan int, 10)
	in2 := make(chan int, 10)
	out := make(chan int, 10)

	defer close(in1)
	defer close(in2)

	fn := func(i int) int { return i * i }
	go merge2Channels(fn, in1, in2, out, numJobs)

	for b := 1; b <= numJobs; b++ {
		in1 <- b
		in2 <- b + 1
	}

	for v := range out {
		fmt.Printf("Result = %v\n", v)
	}

	// Лабораторная
	portScanner.Scan(1, 10000, 4, 5)
}

func increment(ch chan int, n int) {
	ch <- n + 1
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	result := make(chan int, 10)

	select {
	case val := <-firstChan:
		result <- val * val
	case val := <-secondChan:
		result <- val * 3
	case <-stopChan:
		close(result)
	}

	return result
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {

	for b := 1; b <= n; b++ {
		x1 := <-in1
		x2 := <-in2
		out <- fn(x1) + fn(x2)
	}
	close(out)
}

type Str struct{}
