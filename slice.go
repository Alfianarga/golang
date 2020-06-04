package main

import "fmt"

func main() {
	slice1 := []int{1,2}
	slice2 := append(slice1, 3)
	fmt.Println(slice1, slice2)

	slice3 := []int{1,2,3}
	slice4 := make([]int, 3)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}