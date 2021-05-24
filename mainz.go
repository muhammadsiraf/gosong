package main

import (
	"fmt"
	"strings"
)

func main() {
	var isExist = strings.Count("Jasmine Van J", "J")
	fmt.Println(isExist)
	var text = "banana"
	var find = "a"
	var replace = "o"
	var indek = strings.Replace(text, find, replace, 3)
	fmt.Println(indek)
	var zikir = strings.Repeat("Lailahailalah ", 99)
	fmt.Println(zikir)
	var kata = strings.Split("mama mia lexatos", " ")
	fmt.Println(kata[0])

	var data = []string{"banana", "papaya", "tomato"}
	var str = strings.Join(data, "-")
}
