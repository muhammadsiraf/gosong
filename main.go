package main

import (
	"fmt"
	"math"
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

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jarijari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jarijari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

type Hewan struct {
	Nama    string
	Spesies string
}

func main() {
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("Persegi")
	fmt.Println("luas    ", bangunDatar.luas())
	fmt.Println("keliling    ", bangunDatar.keliling())

	bangunDatar = lingkaran{32.1}
	fmt.Println("lingkaran")
	fmt.Println("luas   ", bangunDatar.luas())
	fmt.Println("keliling", bangunDatar.keliling())
	fmt.Println("keliling", bangunDatar.(lingkaran).jarijari())

	var kucing interface{}
	kucing = Hewan{"kucing", "Felis Felis"}
	fmt.Println(kucing.(Hewan).Nama)
	fmt.Println(kucing.(Hewan).Spesies)

	var bangunRuang hitungV2 = &kubus{4}
	fmt.Println(bangunRuang.luas())
	fmt.Println(bangunRuang.keliling())
	fmt.Println(bangunRuang.volume())

}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitungV2 interface {
	hitung2d
	hitung3d
}
