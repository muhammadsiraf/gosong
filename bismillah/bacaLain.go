package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	i, _ := strconv.Atoi(scanner.Text())
	fmt.Println(i)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
