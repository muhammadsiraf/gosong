package main

import "fmt"

func main() {
	orderSomeFood("pizza")
	orderSomeFood("pecel")
	orderSomeFood("burger")
	orderSomeFood("rujak")
}

func orderSomeFood(menu string) {
	defer fmt.Println("Terima Kasih preketek")
	if menu == "pecel" {
		fmt.Print("Pilihan tepat", " ")
		fmt.Print("Puecel e enak broq!", "\n")
		return
	}

	fmt.Println("Pesanan anda:", menu)
}
