package main

import (
	"fmt"

	"github.com/muhammadsiraf/gosong/lib"
)

func main() {

	var s1 = lib.Student{"Udin", 23}
	fmt.Println("name", s1.Name)
	fmt.Println("name", s1.Grade)
}
