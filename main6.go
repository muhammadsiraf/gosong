package main

import (
	"fmt"
	"runtime"
)

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}
	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	ch <- max
}

func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
}

func printMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	go sendMessage(messages)
	printMessage(messages)
}
