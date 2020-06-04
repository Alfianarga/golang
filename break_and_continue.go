package main

import "fmt"

func main()  {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Memberhentikan perulangan")
			break // break disini
		}
		fmt.Println("Nilai dari variable i adalah", i)
	}
	fmt.Println("Keluar Program break")
	println("==========\n")

	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Memberhentikan perulangan")
			continue // continue disini
		}
		fmt.Println("Nilai dari variable i adalah", i)
	}
	fmt.Println("Keluar Program continue")
}
