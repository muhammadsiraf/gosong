package lib

import "fmt"

var ikan string

func init() {
	ikan = "Lele"
}

func Barakuda() {
	fmt.Println("Barakuda", ikan)
}
