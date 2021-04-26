package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var number = 23
	// var reflectValue = reflect.ValueOf(number)
	// fmt.Println("tipe variable :", reflectValue.Type())

	// if reflectValue.Kind() == reflect.Int {
	// 	fmt.Println("nilai variabel :", reflectValue.Int())
	// }
	// haha := reflectValue.Interface().(int)
	// fmt.Println("nilai varibales : ", reflectValue.Interface().(int))
	// fmt.Println(haha)

	var s1 = &student{Name: "Wick", Grade: 2}
	s1.getPropertyInfo()

	var s2 = &student{
		Name:  "Hojn CIkw",
		Grade: 3,
	}
	fmt.Println("nama :", s2.Name)

	var reflectValue = reflect.ValueOf(s2)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Wick"),
	})

	fmt.Println("nama :", s2.Name)
}

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)
	fmt.Println(reflectValue)
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama		:", reflectType.Field(i).Name)
		fmt.Println("tipe data  :", reflectType.Field(i).Type)
		fmt.Println("nilai		:", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}
