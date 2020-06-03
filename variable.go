package main

import "fmt"

func main() {
	var x string = "Hello World"
	fmt.Println(x)

	var y string
	y = "Hello World"
	fmt.Println(y)

	var a string
	a = "first"
	fmt.Println(a)
	a = "second"
	fmt.Println(a)

	var b string
	b = "first "
	fmt.Println(b)
	b = b + "second"
	fmt.Println(b)

	var t string = "hello"
	var r string = "world"
	fmt.Println(t == r)

	u := "hello"
	fmt.Println(u)

	i := 1
	fmt.Println(i)

	name := "Max"
	fmt.Println("My dog's name is", name)

	const o string = "Hello World"
	fmt.Println(o)

	//multiple variable
	var (
		l = 5
		k = 10
		j = 15
	)
	total := l + k + j
	fmt.Println(l, "+", k, "+", j, "=", total)
}
