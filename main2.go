package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA
	fmt.Println(*numberB)

	var numberZ int = 5
	var numberX *int = &numberZ

	fmt.Println(numberZ)
	fmt.Println(&numberZ)

	fmt.Println(numberX)
	fmt.Println(*numberX)

	var numberC **int = &numberX
	fmt.Println(**numberC)
	fmt.Println(numberC)

}
