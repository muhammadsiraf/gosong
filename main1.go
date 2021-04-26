package main

import (
	"fmt"
	"strings"
)

func main() {
	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "mango", "bango"}
	fmt.Println(secret)

	secret = 12.3
	fmt.Println(secret)

	var data map[string]interface{}

	data = map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"braskfast": []string{"apple", "mango", "banana"},
	}

	fmt.Println(data)

	data2 := make(map[string]string)
	data2["sour"] = "asem"
	data2["sweet"] = "legi"
	data2["bitter"] = "pait"
	data2["spicy"] = "pedes"
	data2["salty"] = "asin"

	fmt.Println(data2)

	var secret2 interface{}

	secret2 = 2
	var number = secret2.(int) * 10
	fmt.Println(secret2, "muliplied by 10 is :", number)

	secret2 = []string{"apple", "mango", "banana"}
	var gruits = strings.Join(secret2.([]string), ", ")
	fmt.Println(gruits, "is my favourite fruits")
}
