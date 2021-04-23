package lib

import "fmt"

type Student struct {
	Name  string
	Grade int
}

func SayHello() {
	fmt.Println("Bonjour")
	introduceMe("Udin")
}

func SayGoodBye() {
	fmt.Println("Au Revoir")
}

func introduceMe(name string) {
	fmt.Println("Bonjour my name is ", name)
}
