package main

import (
	"fmt"

	"github.com/muhammadsiraf/gosong/lib"
)

var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "John Wick"
	Student.Grade = 3
	fmt.Println("guide us help us")
}

func main() {
	lib.Barakuda()
}
