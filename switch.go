package main

import "fmt"

func main() {
	fmt.Print("Masukan Angka dari 1 - 3 = ")
	var input int64
	fmt.Scanf("%d", &input)

	switch input {
	case 1:
		fmt.Println("Kamu memilih angka 1")
	case 2:
		fmt.Println("Kamu memilih angka 2")
	case 3:
		fmt.Println("Kamu memilih angka 3")
	default:
		fmt.Println("Silahkan masukan angka 1-3")
	}
}
