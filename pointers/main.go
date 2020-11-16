package main

import (
	"fmt"
)

var a int = 10
var b *int = &a
var c int = *b

func main() {
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)

	fmt.Println("b = ", b)
	fmt.Println("*b = ", *b)

	fmt.Println("c = ", c)

	fmt.Println("*&a = ", *&a)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
}
