package main

import "fmt"

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "wick",
	height:      182.3,
	age:         29,
	isGraduated: false,
	hobbies:     []string{"eating", "sleeping"},
}

func main() {
	fmt.Printf("%d\n", data.age)
	fmt.Printf("%.1f\n", data.height)
	fmt.Printf("%g\n", 0.122222222222222222222222)
	fmt.Printf("%o\n", data.age)
	fmt.Printf("%p\n", &data.age)
	fmt.Printf("%q\n", `" name \ height "`)
	fmt.Printf("%s\n", data.name)
	fmt.Printf("%t\n", data.isGraduated)
	fmt.Printf("%T\n", data.name)
	fmt.Printf("%v\n", data.name)
	fmt.Printf("%v\n", data.height)
	fmt.Printf("%+v\n", data)
	fmt.Printf("%#v\n", data)
	fmt.Printf("%x\n", data.age)

	var d = data.name

	fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3])

	fmt.Printf("%x\n", d)
	fmt.Printf("%%\n")

}
