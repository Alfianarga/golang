package main

import "fmt"

func main() {
	mobil := [3]string{"avanza", "xenia", "bmw"}
	fmt.Println(mobil[1])

	var x [5]int
	x[2]=100
	fmt.Println("dari panjang array berisi",len(x),"mempunyai data sebagai berikut",x)
}
