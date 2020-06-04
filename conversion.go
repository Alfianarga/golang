package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(reflect.TypeOf(s).String())
	fmt.Println(reflect.TypeOf(i).String())
}
