package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go process("order")
	go func() {
		process("refund")
		wg.Done()
	}()
	wg.Wait()
	// fmt.Scanln()
}

func process(item string, out chan string) {
	for i := 1; 1 <= 5; i++ {
		time.Sleep(time.Second / 2)
		fmt.Println("Processed", i, item)
	}
}
