package main

import "fmt"

func main() {
	//Numbers
	fmt.Println("1 + 1 = ",1 + 1)
	fmt.Println("1 + 1 = ",1.0 + 1.0)

	//String
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello World"[1:4])
	fmt.Println("Hello World"[:1])
	fmt.Println("Hello World"[2:])
	fmt.Println("Hello " + "World")

	//Boolean
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}